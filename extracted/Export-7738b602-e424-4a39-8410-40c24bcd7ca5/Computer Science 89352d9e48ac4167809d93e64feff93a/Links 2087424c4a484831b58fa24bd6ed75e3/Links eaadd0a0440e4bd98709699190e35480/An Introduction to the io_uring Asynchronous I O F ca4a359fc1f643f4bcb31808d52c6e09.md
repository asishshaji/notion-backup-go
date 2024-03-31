# An Introduction to the io_uring Asynchronous I/O Framework

URL: https://medium.com/oracledevs/an-introduction-to-the-io-uring-asynchronous-i-o-framework-fad002d7dfc1
Category: Linux, Optimisation

![An%20Introduction%20to%20the%20io_uring%20Asynchronous%20I%20O%20F%20ca4a359fc1f643f4bcb31808d52c6e09/0TEuhyiZGjbXQ8zbD.png](An%20Introduction%20to%20the%20io_uring%20Asynchronous%20I%20O%20F%20ca4a359fc1f643f4bcb31808d52c6e09/0TEuhyiZGjbXQ8zbD.png)

*In this blog Oracle Linux kernel developer **Bijan Mottahedeh** talks about the io_uring asynchronous I/O framework included in the Unbreakable Enterprise Kernel 6.*

# Introduction

This blog post is a brief introduction to the io_uring asynchronous I/O framework available in release 6 of the Unbreakable Enterprise Kernel (UEK). It highlights the motivations for introducing the new framework, describes its system call and library interfaces with the help of sample applications, and provides a list of references that describe the technology in further detail including more usage examples.

The io_uring Asynchronous I/O (AIO) framework is a new Linux I/O interface, first introduced in the upstream Linux kernel version 5.1 (March 2019). It provides a low-latency and feature-rich interface for applications that require AIO functionality but prefer the kernel to perform the I/O. This could be in order to exploit benefits running on top of a filesystem or to leverage features such as mirroring and block-level encryption. This is in contrast to SPDK applications for example, that explicitly do not want the kernel to perform I/O because they implement their own filesystem and features.

# Motivation

The native Linux AIO framework suffers from various limitations, which io_uring aims to overcome:

- It does not support buffered I/O, only direct I/O is supported.
- It has non-deterministic behavior which may block under various circumstances.
- It has a sub-optimal API, which requires at least two system calls per I/O, one to submit a request, and one to wait for its completion.

# Communication Channel

An io_uring instance has two rings, a submission queue (SQ) and a completion queue (CQ), shared between the kernel and the application. The queues are single producer, single consumer, and power of two in size.

The queues provide a lock-less access interface, coordinated with memory barriers.

The application creates one or more SQ entries (SQE), and then updates the SQ tail. The kernel consumes the SQEs , and updates the SQ head.

The kernel creates CQ entries (CQE) for one or more completed requests, and updates the CQ tail. The application consumes the CQEs and updates the CQ head.

Completion events may arrive in any order but they are always associated with specific SQEs.

# System call API

The io_uring API consists of three system calls: io_uring_setup(2), io_uring_register(2) and io_uring_enter(2), described in the following sections. The full manual pages for the system calls are available [here](https://github.com/axboe/liburing/tree/master/man).

# io_uring_setup

*set up a context for performing asynchronous I/O*

```
int io_uring_setup(u32 entries, struct io_uring_params *p);
```

The io_uring_setup() system call sets up a submission queue and completion queue with at least *entries* elements, and returns a file descriptor which can be used to perform subsequent operations on the io_uring instance. The submission and completion queues are shared between the application and the kernel, which eliminates the need to copy data when initiating and completing I/O.

*params* is used by the application in order to configure the io_uring instance, and by the kernel to convey back information about the configured submission and completion ring buffers.

An io_uring instance can be configured in three main operating modes:

- **Interrupt driven** — By default, the io_uring instance is setup for interrupt driven I/O. I/O may be submitted using io_uring_enter() and can be reaped by checking the completion queue directly.
- **Polled** — Perform busy-waiting for an I/O completion, as opposed to getting notifications via an asynchronous IRQ (Interrupt Request). The file system (if any) and block device must support polling in order for this to work. Busy-waiting provides lower latency, but may consume more CPU resources than interrupt driven I/O. Currently, this feature is usable only on a file descriptor opened using the *O_DIRECT* flag. When a read or write is submitted to a polled context, the application must poll for completions on the CQ ring by calling io_uring_enter(). It is illegal to mix and match polled and non-polled I/O on an io_uring instance.
- **Kernel polled** — In this mode, a kernel thread is created to perform submission queue polling. An io_uring instance configured in this way enables an application to issue I/O without ever context switching into the kernel. By using the submission queue to fill in new submission queue entries and watching for completions on the completion queue, the application can submit and reap I/Os without doing a single system call. If the kernel thread is idle for more than a user configurable amount of time, it will go idle after notifying the application first. When this happens, the application must call io_uring_enter() to wake the kernel thread. If I/O is kept busy, the kernel thread will never sleep.

io_uring_setup() returns a new file descriptor on success. The application may then provide the file descriptor in a subsequent mmap(2) call to map the submission and completion queues, or to the io_uring_register() or io_uring_enter() system calls.

# io_uring_register

*register files or user buffers for asynchronous I/O*

```
int io_uring_register(unsigned int fd, unsigned int opcode, void *arg, unsigned int nr_args);
```

The io_uring_register() system call registers user buffers or files for use in an io_uring instance referenced by *fd*. Registering files or user buffers allows the kernel to take long term references to internal kernel data structures associated with the files, or create long term mappings of application memory associated with the buffers, only once during registration as opposed to during processing of each I/O request, therefore reducing per-I/O overhead.

Registered buffers will be locked in memory and charged against the user’s *RLIMIT_MEMLOCK* resource limit. Additionally, there is a size limit of 1GiB per buffer. Currently, the buffers must be anonymous, non-file-backed memory, such as that returned by malloc(3) or mmap(2) with the *MAP_ANONYMOUS* flag set. Huge pages are supported as well. Note that the entire huge page will be pinned in the kernel, even if only a portion of it is used.

It is perfectly valid to setup a large buffer and then only use part of it for an I/O, as long as the range is within the originally mapped region.

An application can increase or decrease the size or number of registered buffers by first unregistering the existing buffers, and then issuing a new call to io_uring_register() with the new buffers.

An application can dynamically update the set of registered files without unregistering them first.

It is possible to use eventfd(2) to get notified of completion events on an io_uring instance. If this is desired, an eventfd file descriptor can be registered through this system call.

The credentials of the running application can be registered with io_uring which returns an id associated with those credentials. Applications wishing to share a ring between separate users/processes can pass in this credential id in the SQE personality field. If set, that particular SQE will be issued with these credentials.

# io_uring_enter

*initiate and/or complete asynchronous I/O*

```
int io_uring_enter(unsigned int fd, unsigned int to_submit, unsigned int min_complete, unsigned int flags, sigset_t *sig);
```

io_uring_enter() is used to initiate and complete I/O using the shared submission and completion queues setup by a call to io_uring_setup(). A single call can both submit new I/O and wait for completions of I/O initiated by this call or previous calls to io_uring_enter().

*fd* is the file descriptor returned by io_uring_setup(). *to_submit* specifies the number of I/Os to submit from the submission queue. If so directed by the application, the system call will attempt to wait for *min_complete* event completions before returning. If the io_uring instance was configured for polling, then *min_complete* has a slightly different meaning. Passing a value of 0 instructs the kernel to return any events which are already complete, without blocking. If *min_complete* is a non-zero value, the kernel will still return immediately if any completion events are available. If no event completions are available, then the call will poll either until one or more completions become available, or until the process has exceeded its scheduler time slice.

Note that for interrupt driven I/O, an application may check the completion queue for event completions without entering the kernel at all.

io_uring_enter() supports a wide variety of operations, including

- Open, close, and stat files
- Read and write into multiple buffers or pre-mapped buffers
- Socket I/O operations
- Synchronize file state
- Asynchronously monitor a set of file descriptors
- Create a timeout linked to a specific operation in the ring
- Attempt to cancel an operation that is currently in flight
- Create I/O chains
- Ordered execution within a chain
- Parallel execution of multiple chains

When the system call returns that a certain amount of SQEs have been consumed and submitted, it’s safe to reuse SQE entries in the ring. This is true even if the actual IO submission had to be punted to async context, which means that the SQE may in fact not have been submitted yet. If the kernel requires later use of a particular SQE entry, it will have made a private copy of it.

# Liburing

Liburing provides a simple higher level API for basic use cases and allows applications to avoid having to deal with the full system call implementation details. The API also avoids duplicated code for operations such as setting up an io_uring instance.

For example, after getting back a ring file descriptor from io_uring_setup(), an application must always call mmap() in order to map the submission and completion queues for access, as described in the io_uring_setup() manual page. The entire sequence is somewhat lengthy but can be accomplish with the following liburing call:

```
int io_uring_queue_init(unsigned entries, struct io_uring *ring, unsigned flags);
```

The sample applications below are included in liburing source, and help illustrates these points.

The source code for liburing sample applications is available [here](https://github.com/axboe/liburing/tree/master/examples).

There is no liburing API documentation currently available and the API is described in [liburing.h](https://github.com/axboe/liburing/blob/master/src/include/liburing.h) header file.

# Sample application io_uring-test

io_uring-test reads a maximum of 16KB from a user specified file using 4 SQEs. Each SQE is a request to read 4KB of data from a fixed file offset. io-uring then reaps each CQE and checks whether the full 4KB was read from the file as requested.

If the file is smaller than 16KB, all 4 SQEs are still submitted but some CQE results will indicate either a partial read, or zero bytes read, depending on the actual size of the file.

io-uring finally reports the number of SQEs and CQEs it has processed.

The full source code is listed below followed by the description of the liburing calls.

```
/* SPDX-License-Identifier: MIT */
/*
 * Simple app that demonstrates how to setup an io_uring interface,
 * submit and complete IO against it, and then tear it down.
 *
 * gcc -Wall -O2 -D_GNU_SOURCE -o io_uring-test io_uring-test.c -luring
 */
#include <stdio.h>
#include <fcntl.h>
#include <string.h>
#include <stdlib.h>
#include <unistd.h>
#include "liburing.h"

#define QD  4

int main(int argc, char *argv[])
{
    struct io_uring ring;
    int i, fd, ret, pending, done;
    struct io_uring_sqe *sqe;
    struct io_uring_cqe *cqe;
    struct iovec *iovecs;
    off_t offset;
    void *buf;

    if (argc < 2) {
        printf("%s: file\n", argv[0]);
        return 1;
    }

    ret = io_uring_queue_init(QD, &ring, 0);
    if (ret < 0) {
        fprintf(stderr, "queue_init: %s\n", strerror(-ret));
        return 1;
    }

    fd = open(argv[1], O_RDONLY | O_DIRECT);
    if (fd < 0) {
        perror("open");
        return 1;
    }

    iovecs = calloc(QD, sizeof(struct iovec));
    for (i = 0; i < QD; i++) {
        if (posix_memalign(&buf, 4096, 4096))
            return 1;
        iovecs[i].iov_base = buf;
        iovecs[i].iov_len = 4096;
    }

    offset = 0;
    i = 0;
    do {
        sqe = io_uring_get_sqe(&ring);
        if (!sqe)
            break;
        io_uring_prep_readv(sqe, fd, &iovecs[i], 1, offset);
        offset += iovecs[i].iov_len;
        i++;
    } while (1);

    ret = io_uring_submit(&ring);
    if (ret < 0) {
        fprintf(stderr, "io_uring_submit: %s\n", strerror(-ret));
        return 1;
    }

    done = 0;
    pending = ret;
    for (i = 0; i < pending; i++) {
        ret = io_uring_wait_cqe(&ring, &cqe);
        if (ret < 0) {
            fprintf(stderr, "io_uring_wait_cqe: %s\n", strerror(-ret));
            return 1;
        }

        done++;
        ret = 0;
        if (cqe->res != 4096) {
            fprintf(stderr, "ret=%d, wanted 4096\n", cqe->res);
            ret = 1;
        }
        io_uring_cqe_seen(&ring, cqe);
        if (ret)
            break;
    }

    printf("Submitted=%d, completed=%d\n", pending, done);
    close(fd);
    io_uring_queue_exit(&ring);
    return 0;
}
```

# Description

An io_uring instance is created in the default interrupt driven mode, specifying only the size of the ring.

```
ret = io_uring_queue_init(QD, &ring, 0);
```

All of the ring SQEs are next fetched and prepared for the IORING_OP_READV operation which provides an asynchronous interface to readv(2) system call. Liburing provides numerous helper functions to prepare io_uring operations.

Each SQE will point to an allocated buffer described by an iovec structure. The buffer will contain the result of the corresponding readv operation upon completion.

```
sqe = io_uring_get_sqe(&ring); io_uring_prep_readv(sqe, fd, &iovecs[i], 1, offset);
```

The SQEs are submitted with a single call to io_uring_submit() which returns the number of submitted SQEs.

```
ret = io_uring_submit(&ring);
```

The CQEs are reaped with repeated calls to io_uring_wait_cqe(), and the success of a given submission is verified with the cqe->res field; each matching call to io_uring_cqe_seen() informs the kernel that the given CQE has been consumed.

```
ret = io_uring_wait_cqe(&ring, &cqe); io_uring_cqe_seen(&ring, cqe);
```

The io_uring instance is finally dismantled.

```
void io_uring_queue_exit(struct io_uring *ring)
```

# Sample application link-cp

link-cp copies a file using the io_uring SQE chaining feature.

As noted before, io_uring supports the creation of I/O chains. The I/O operations within a chain are sequentially executed, and multiple I/O chains can execute in parallel.

To copy the file, link-cp creates SQE chains of length two. The first SQE in the chain is a read request from the input file into a buffer. The second request, linked to the first, is a write request from the same buffer to the output file.

```
/* SPDX-License-Identifier: MIT */
/*
 * Very basic proof-of-concept for doing a copy with linked SQEs. Needs a
 * bit of error handling and short read love.
 */
#include <stdio.h>
#include <fcntl.h>
#include <string.h>
#include <stdlib.h>
#include <unistd.h>
#include <assert.h>
#include <errno.h>
#include <inttypes.h>
#include <sys/types.h>
#include <sys/stat.h>
#include <sys/ioctl.h>
#include "liburing.h"

#define QD  64
#define BS  (32*1024)

struct io_data {
    size_t offset;
    int index;
    struct iovec iov;
};

static int infd, outfd;
static unsigned inflight;

static int setup_context(unsigned entries, struct io_uring *ring)
{
    int ret;

    ret = io_uring_queue_init(entries, ring, 0);
    if (ret < 0) {
        fprintf(stderr, "queue_init: %s\n", strerror(-ret));
        return -1;
    }

    return 0;
}

static int get_file_size(int fd, off_t *size)
{
    struct stat st;

    if (fstat(fd, &st) < 0)
        return -1;
    if (S_ISREG(st.st_mode)) {
        *size = st.st_size;
        return 0;
    } else if (S_ISBLK(st.st_mode)) {
        unsigned long long bytes;

        if (ioctl(fd, BLKGETSIZE64, &bytes) != 0)
            return -1;

        *size = bytes;
        return 0;
    }

    return -1;
}

static void queue_rw_pair(struct io_uring *ring, off_t size, off_t offset)
{
    struct io_uring_sqe *sqe;
    struct io_data *data;
    void *ptr;

    ptr = malloc(size + sizeof(*data));
    data = ptr + size;
    data->index = 0;
    data->offset = offset;
    data->iov.iov_base = ptr;
    data->iov.iov_len = size;

    sqe = io_uring_get_sqe(ring);
    io_uring_prep_readv(sqe, infd, &data->iov, 1, offset);
    sqe->flags |= IOSQE_IO_LINK;
    io_uring_sqe_set_data(sqe, data);

    sqe = io_uring_get_sqe(ring);
    io_uring_prep_writev(sqe, outfd, &data->iov, 1, offset);
    io_uring_sqe_set_data(sqe, data);
}

static int handle_cqe(struct io_uring *ring, struct io_uring_cqe *cqe)
{
    struct io_data *data = io_uring_cqe_get_data(cqe);
    int ret = 0;

    data->index++;

    if (cqe->res < 0) {
        if (cqe->res == -ECANCELED) {
            queue_rw_pair(ring, BS, data->offset);
            inflight += 2;
        } else {
            printf("cqe error: %s\n", strerror(cqe->res));
            ret = 1;
        }
    }

    if (data->index == 2) {
        void *ptr = (void *) data - data->iov.iov_len;

        free(ptr);
    }
    io_uring_cqe_seen(ring, cqe);
    return ret;
}

static int copy_file(struct io_uring *ring, off_t insize)
{
    struct io_uring_cqe *cqe;
    size_t this_size;
    off_t offset;

    offset = 0;
    while (insize) {
        int has_inflight = inflight;
        int depth;

        while (insize && inflight < QD) {
            this_size = BS;
            if (this_size > insize)
                this_size = insize;
            queue_rw_pair(ring, this_size, offset);
            offset += this_size;
            insize -= this_size;
            inflight += 2;
        }

        if (has_inflight != inflight)
            io_uring_submit(ring);

        if (insize)
            depth = QD;
        else
            depth = 1;
        while (inflight >= depth) {
            int ret;

            ret = io_uring_wait_cqe(ring, &cqe);
            if (ret < 0) {
                printf("wait cqe: %s\n", strerror(ret));
                return 1;
            }
            if (handle_cqe(ring, cqe))
                return 1;
            inflight--;
        }
    }

    return 0;
}

int main(int argc, char *argv[])
{
    struct io_uring ring;
    off_t insize;
    int ret;

    if (argc < 3) {
        printf("%s: infile outfile\n", argv[0]);
        return 1;
    }

    infd = open(argv[1], O_RDONLY);
    if (infd < 0) {
        perror("open infile");
        return 1;
    }
    outfd = open(argv[2], O_WRONLY | O_CREAT | O_TRUNC, 0644);
    if (outfd < 0) {
        perror("open outfile");
        return 1;
    }

    if (setup_context(QD, &ring))
        return 1;
    if (get_file_size(infd, &insize))
        return 1;

    ret = copy_file(&ring, insize);

    close(infd);
    close(outfd);
    io_uring_queue_exit(&ring);
    return ret;
}
```

# Description

The three routines copy_file(), queue_rw_pair(), and handle_cqe(), implement the file copy.

copy_file() implements the high level copy loop; it calls queue_rw_pair() to construct each SQE pair

```
queue_rw_pair(ring, this_size, offset);
```

and submits all the constructed SQE pairs in each iteration with a single call to io_uring_submit().

```
if (has_inflight != inflight) io_uring_submit(ring);
```

copy_file() maintains up to QD SQEs inflight as long as data remains to be copied; it waits for and reaps all CQEs after the input file has been fully read.

```
ret = io_uring_wait_cqe(ring, &cqe); if (handle_cqe(ring, cqe))
```

queue_rw_pair() constructs a read-write SQE pair. The IOSQE_IO_LINK flag is set for the read SQE which marks the start of the chain. The flag is not set for the write SQE which marks the end of the chain. The user data field is set for both SQEs to the same data descriptor for the pair, and will be used during completion handling.

```
sqe = io_uring_get_sqe(ring); io_uring_prep_readv(sqe, infd, &data->iov, 1, offset); sqe->flags |= IOSQE_IO_LINK; io_uring_sqe_set_data(sqe, data); sqe = io_uring_get_sqe(ring); io_uring_prep_writev(sqe, outfd, &data->iov, 1, offset); io_uring_sqe_set_data(sqe, data);
```

handle_cqe() retrieves from a CQE the data descriptor saved by queue_rw_pair() and records the retrieval in the descriptor.

```
struct io_data *data = io_uring_cqe_get_data(cqe); data->index++;
```

handle_cqe() resubmits the read-write pair if the request was canceled.

```
if (cqe->res == -ECANCELED) { queue_rw_pair(ring, BS, data->offset);
```

The following excerpt from the io_uring_enter() manual page describes the behavior of chained requests in more detail:

# **IOSQE_IO_LINK**

When this flag is specified, it forms a link with the next SQE in the submission ring. That next SQE will not be started before this one completes. This, in effect, forms a chain of SQEs, which can be arbitrarily long. The tail of the chain is denoted by the first SQE that does not have this flag set. This flag has no effect on previous SQE submissions, nor does it impact SQEs that are outside of the chain tail. This means that multiple chains can be executing in parallel, or chains and individual SQEs. Only members inside the chain are serialized. A chain of SQEs will be broken, if any request in that chain ends in error. io_uring considers any unexpected result an error. This means that, eg, a short read will also terminate the remainder of the chain. If a chain of SQE links is broken, the remaining unstarted part of the chain will be terminated and completed with -ECANCELED as the error code.

handle_cqe() frees the shared data descriptor after both members of a CQE pair have been processed.

```
if (data->index == 2) { void *ptr = (void *) data - data->iov.iov_len; free(ptr); }
```

handle_cqe() finally informs the kernel that the given CQE has been consumed.

```
io_uring_cqe_seen(ring, cqe);
```

# Liburing API

io_uring-test and link-cp use the following subset of the liburing API:

```
/*
 * Returns -1 on error, or zero on success. On success, 'ring'
 * contains the necessary information to read/write to the rings.
 */
int io_uring_queue_init(unsigned entries, struct io_uring *ring, unsigned flags);

/*
 * Return an sqe to fill. Application must later call io_uring_submit()
 * when it's ready to tell the kernel about it. The caller may call this
 * function multiple times before calling io_uring_submit().
 *
 * Returns a vacant sqe, or NULL if we're full.
 */
struct io_uring_sqe *io_uring_get_sqe(struct io_uring *ring);

/*
 * Set the SQE user_data field.
 */
void io_uring_sqe_set_data(struct io_uring_sqe *sqe, void *data);

/*
 * Prepare a readv I/O operation.
 */
void io_uring_prep_readv(struct io_uring_sqe *sqe, int fd,
                         const struct iovec *iovecs,
                         unsigned nr_vecs, off_t offset);

/*
 * Prepare a writev I/O operation.
 */
void io_uring_prep_writev(struct io_uring_sqe *sqe, int fd,
                          const struct iovec *iovecs,
                          unsigned nr_vecs, off_t offset);

/*
 * Submit sqes acquired from io_uring_get_sqe() to the kernel.
 *
 * Returns number of sqes submitted
 */
int io_uring_submit(struct io_uring *ring);

/*
 * Return an IO completion, waiting for it if necessary. Returns 0 with
 * cqe_ptr filled in on success, -errno on failure.
 */
int io_uring_wait_cqe(struct io_uring *ring,
                      struct io_uring_cqe **cqe_ptr);

/*
 * Must be called after io_uring_{peek,wait}_cqe() after the cqe has
 * been processed by the application.
 */
static inline void io_uring_cqe_seen(struct io_uring *ring,
                                     struct io_uring_cqe *cqe);

void io_uring_queue_exit(struct io_uring *ring);
```

# References

- [Efficient IO with io_uring](https://kernel.dk/io_uring.pdf)
- [Ringing in a new asynchronous I/O API](https://lwn.net/Articles/776703/)
- [The rapid growth of io_uring](https://lwn.net/Articles/810414/)
- [Faster I/O through io_uring](https://www.youtube.com/watch?v=-5T4Cjw46ys)
- [System call API](https://github.com/axboe/liburing/tree/master/man)
- [Liburing API](https://github.com/axboe/liburing/blob/master/src/include/liburing.h)
- [Examples](https://github.com/axboe/liburing/tree/master/examples)
- [Liburing repository](https://github.com/axboe/liburing)

*Originally published at [https://blogs.oracle.com](https://blogs.oracle.com/linux/an-introduction-to-the-io_uring-asynchronous-io-framework).*