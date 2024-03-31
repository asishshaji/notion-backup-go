# A Deep-dive into the Linux Filesystem | Part 1

URL: https://emmanuelbashorun.medium.com/a-deep-dive-into-the-linux-filesystem-part-1-d06d365d42bb
Category: Linux, Optimisation

This article is one out of a series of articles that seeks to embark on a low to deep level dive geared towards understanding the Linux filesystem.

This article is cool for:

1. new Linux users wanting to understand the filesystem mechanism,
2. mid-level Linux users that want to have a deep but not academic-type understanding of the filesystem,
3. programmers interested in how the kernel (some kernel dev’ish things) handles files and directories.
4. Or if you’re a professional knowledge hoarder like me

During the course of this article, there’d be some C code snippets around but you won’t need to be a C ninja to understand as I’m not even one. But I’d make sure to have simple explanations for them.

But before throwing lots of Linux filesystem ‘gibberish’ at you, its pertinent to genuinely have an understanding on the working internals of a ‘filesystem’ regardless of the operating system so as the ‘gibberish’ can make some sense.

So, to business…

# What’s a filesystem?

You see a filesystem basically handles how data (files, directories, block devices) is arranged, stored. It could also be described as a collection of these datastructures within the operating system. It can even be seen as the whole logic and program structure that controls the behavior of data.

Wikipedia says a filesystem

> ‘controls how data is stored and retrieved. Without a file system, data placed in a storage medium would be one large body of data with no way to tell where one piece of data stops and the next begins.’ https://en.wikipedia.org/wiki/File_system
> 

In essence the filesystem arranges data in an intelligent way. And the storage medium can be a HDD, SDD, flash drive or even a RAM. There are lots of filesystem implementations (Minix, EXT’s, XFS, ZFS, NTFS etc) around (I’d try to talk a bit about them later) with an individual and one collective (VFS comes later on) method of working with data. A filesystem could even be seen as a filesystem implementation.

The filesystem gives the rules on how data should be accessed on our storage drive. The drive is usually formatted with a type of filesystem implementation to make it accessible to the operating system.

# Filesystem Structure

NB: Data on disks are usually accessed in blocks rather than bytes.

As said earlier, data is organized on disk drives by the filesystems and the organization can be understood in layers.

Firstly, there’s the lowest layers which is the **physical device** or disk drive.

Secondly, the **[Input/Output Control](https://www.cs.uic.edu/~jbell/CourseNotes/OperatingSystems/12_FileSystemImplementation.html)** are the hardware drivers or programs which communicate with the devices by reading and writing special codes directly to and from memory addresses corresponding to the controller card’s register.

Thirdly, the **basic file system** cooperates with the device driver to retrieve and store information of blocks of data without knowing the blocks content.

Fourthly, the **file organization module** has information on how logical blocks are mapped to physical blocks on the disk for data storage or retrieval.

And lastly there’s the **logical file system** handles information about a file i.e. it meta data except the files actual data. Things like superblocks and inodes are handled here.

Figure 1: Filesystem structure

![A%20Deep-dive%20into%20the%20Linux%20Filesystem%20Part%201%206022953d896841e98e628eb9f63ce3d4/1UXwrwlJFa05BzBuGqzbTUA.jpeg](A%20Deep-dive%20into%20the%20Linux%20Filesystem%20Part%201%206022953d896841e98e628eb9f63ce3d4/1UXwrwlJFa05BzBuGqzbTUA.jpeg)

Ref: [https://www.cs.uic.edu/~jbell/CourseNotes/OperatingSystems/12_FileSystemImplementation.html](https://www.cs.uic.edu/~jbell/CourseNotes/OperatingSystems/12_FileSystemImplementation.html)

Data on a disk resides in a partition and a disk can be divided into more partitions.

You create a partition on a disk drive by formatting it with a filesystem type (remember EXT’s, ZFS etc). Partitions are created if you want to independently have multiple drives formatted with different filesystems and operating system all on one physical hard disk.

By the way, a partition within a disk, be it an hard disk or flash drive is divided into sectors. This sector can be about 512 bytes in size. A group of sector is called a block. And a block can be made of one or more sectors. For information to be retrieved from disk, the operating system would have to point to a particular block.

When a partition is created, data is pushed into the first block of the drive which is called the boot block or MBR (Master Boot Record). The boot block contains information that helps bootstrap the machine unto its feet.

The partition also contains the super block which holds information regarding the filesystem’s layout or structure. It also gives information about the drive’s size, permissions and control, how many file resides on the filesystem, inodes location. The superblock stores metatdata of the filesystem.

Check: [https://www.slashroot.in/understanding-file-system-superblock-linux](https://www.slashroot.in/understanding-file-system-superblock-linux)

Another component of a partition are the inodes (**i**ndex **nodes**). Each file is linked to a number called the i-number.

Inodes contain specific information such as owner identifier, type of file, array of pointers to data blocks on disk, file size etc about a particular file-object. They store the attribute of a particular disk block that data resides in.

The inode stores metadata of a file. Feel free to call an inode the superblock of a file. Its similar to how superblocks store metadata of the entire filesystem.

> The inode number indexes a table of inodes in a known location on the device. From the inode number, the kernel’s file system driver can access the inode contents, including the location of the file, thereby allowing access to the file
> 

Ref: [https://en.wikipedia.org/wiki/Inode](https://en.wikipedia.org/wiki/Inode)

Disk space allocation is controlled by inodes. They are created using the **mkfs (1)** command. These [inodes](http://digi-cron.com:8080/filesystems.html) are so special that they can’t even be modified directly but are controlled by user space commands like touch or [systemcalls](https://www.slashroot.in/what-is-system-call-in-unix-and-linux) like **open(2)**, **unlink(2)**. Commands like **chmod(1)** and **chmod(2)** can be used to alter access permissions.

There’s also the data block which contains the actual data on the drive that makes sense to us. This could be your text documents or binary files.

Figure 2: The command shows the existing partition of my hard disk and the size of 1 sector.

![A%20Deep-dive%20into%20the%20Linux%20Filesystem%20Part%201%206022953d896841e98e628eb9f63ce3d4/1M0aOjcezS9Wn_BP5BSlILA.png](A%20Deep-dive%20into%20the%20Linux%20Filesystem%20Part%201%206022953d896841e98e628eb9f63ce3d4/1M0aOjcezS9Wn_BP5BSlILA.png)

sda1 in the image above contains my Windows OS and also has the bootloader. sda5 is a logical partition that was derived from the sda2 extended partition.

There was a time in ‘filesystems history’ that hard disks were designed to only have 4 partitions. This proved to be a problem when sysadmins, programmers and computer users need more than a partition to manage their hard disk. For instance a programmer probably wanted to have a Linux, Windows, FreeBSD, Solaris, MS-DOS and Minix OS all at once on his/her drive. Or may be the programmer needed a separate swap partition rather than a swap file. This couldn’t happen because of the limit imposed by the prevalent design schema of then.

So you see the problem huh?

> To overcome this design problem, extended partitions were invented. This trick allows partitioning a primary partition into sub-partitions. The primary partition thus subdivided is the extended partition; the sub-partitions are logical partitions. They behave like primary partitions, but are created differently. There is no speed difference between them. By using an extended partition you can now have up to 15 partitions per disk.
> 

Ref: [https://www.tldp.org/LDP/sag/html/partitions.html](https://www.tldp.org/LDP/sag/html/partitions.html)

A flash drive on a PC is loaded by device drivers which makes it accessible to the PC but not readable or writable. It is then formatted using a file system of choice to make it usable.

Before using the flash drive or a request to access a file on the flash drive requires Linux to check the superblock to see the formatted filesystem of the drive. Drivers that supports the flash drive’s file system are loaded from the Linux kernel to read the flash drive and place it on a specific directory. The action of placing a block device formatted with a particular file system on a directory is called **mounting** and that place is called the **mount point**.

# A few supported Linux File systems

There are lots of filesystem types supported by the Linux kernel. You can check yours by running the following command on a terminal.

Figure 3: Command shows supported filesystem types

![A%20Deep-dive%20into%20the%20Linux%20Filesystem%20Part%201%206022953d896841e98e628eb9f63ce3d4/12L2oD4vKLEDCbicrUmUxOw.png](A%20Deep-dive%20into%20the%20Linux%20Filesystem%20Part%201%206022953d896841e98e628eb9f63ce3d4/12L2oD4vKLEDCbicrUmUxOw.png)

**Minix**: This filesytem was the defacto filesystem of the Minix OS. Minix had its own problems as it only could access 64MB of storage, wasn’t free and open source

**ext**: This was created by Remy Card a year after Torvald’s debuted the Linux kernel. This filesystem could address up to 2GB of storage and 255-character filenames. It also introduced the Virtual File System abstraction layer.

**ext2**: This came to solve ext’s problem by adding more timestamps to the inodes for file access, inode creation and file modification. It could also access a bigger pool of storage space. This was a proper filesystem

**ext3**: The idea of *Journaling* came with ext3. ext2 could handle problems that results when there’s a power failure during read/write operations on the hard disk. This most of the time leads to great loss of data from related or unrelated files.

> A journaling filesystem is a filesystem that maintains a special file called a journal that is used to repair any inconsistencies that occur as the result of an improper shutdown of a computer. Such shutdowns are usually due to an interruption of the power supply or to a software problem that cannot be resolved without a rebooting.
> 

## [What is a journaling filesystem? — definition by The Linux Information Project (LINFO)](http://www.linfo.org/journaling_filesystem.html)

### [A journaling filesystem is a filesystem that maintains a special file called a journal that is used to repair any…](http://www.linfo.org/journaling_filesystem.html)

[www.linfo.org](http://www.linfo.org/journaling_filesystem.html)

**ext4**: This is an extension of ext3 and also the latest iteration of the EXT’s. It provides a higher performance, resistance to fragmentation, and improved timestamps. It was built to be backward-compatible with ext3 but not vice-versa. [Check here for a solid talk on ext4 features](https://opensource.com/article/18/4/ext4-filesystem).

**ZFS:** This has been touted by people as a next-gen filesystem. It also doubles down as a volume manager. Features include: Copy-on-write, snapshots, cryptographically based check-summing ability (for detecting data corruption). [https://www.freebsd.org/doc/handbook/zfs.htm, l](https://www.freebsd.org/doc/handbook/zfs.html)[https://itsfoss.com/what-is-zfs/](https://itsfoss.com/what-is-zfs/).

Others include [xfs(5)](http://man7.org/linux/man-pages/man5/xfs.5.html); [btrfs](https://btrfs.wiki.kernel.org/index.php/Main_Page) and **man 8 btrfs** for btrfs on the man page.

**Up next!**

The Linux filesystem hierarchy, permissions, access controls, mounting and formatting of partitions would be discussed in a subsequent article.

Thanks and see you later!