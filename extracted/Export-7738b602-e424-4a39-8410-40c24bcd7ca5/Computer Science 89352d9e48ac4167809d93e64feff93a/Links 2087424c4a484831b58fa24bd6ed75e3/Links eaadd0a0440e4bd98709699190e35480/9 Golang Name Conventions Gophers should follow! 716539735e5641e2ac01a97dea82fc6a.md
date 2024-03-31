# 9 Golang Name Conventions Gophers should follow!

URL: https://blog.devgenius.io/golang-name-convention-gophers-should-follow-e4397fba5dce
Category: Golang

![https://miro.medium.com/v2/resize:fit:1200/1*6Y9lMG-6DTpDw7Oc45KyHw.png](https://miro.medium.com/v2/resize:fit:1200/1*6Y9lMG-6DTpDw7Oc45KyHw.png)

---

In this article, I will talk about the Golang name conventions. It is the first and the most essential thing when you decide to become a Gopher from any other programming language.

**[Click to become a medium member and read unlimited stories**!](https://medium.com/@ramseyjiang_22278/membership)

![9%20Golang%20Name%20Conventions%20Gophers%20should%20follow!%20716539735e5641e2ac01a97dea82fc6a/16Y9lMG-6DTpDw7Oc45KyHw.png](9%20Golang%20Name%20Conventions%20Gophers%20should%20follow!%20716539735e5641e2ac01a97dea82fc6a/16Y9lMG-6DTpDw7Oc45KyHw.png)

Nowadays, Golang has become more and more popular, and lots of developers have started to learn Golang. But developers have different backgrounds, such as Python, PHP, Java, etc. That will make an issue, Golang has its own name convention and is a little bit different from other languages. Though a developer knows all name conventions, sometimes perhaps not remember. I forgot some name conventions sometimes also. But I always use a tool to check all my codes before I commit them. Let’s dive into the Golang name conventions first.

# Golang Name Conventions

The Golang name conventions have 3 layers as follows:

> — Package Names
> 
> 
> — — Files Names
> 
> — — — Function
> 
> — — — Structure
> 
> — — — Variable
> 
> — — — Constant
> 

Perhaps someone will say it has another layer named the project name layer. Actually, if you know Golang deeply, **a project is also a package name, the project just includes several packages in it.**

# Package Name Convention

The following words are from go official link [https://go.dev/doc/effective_go#package-names](https://go.dev/doc/effective_go#package-names).

> By convention, package names are written in lowercase and consist of a single word; there shouldn’t be a need for underscores or mixed caps.
> 

However, **a project is a special package.**

**The project name the best way is to use a single word or an abbreviation**.

*Example:*

```
github.com/username/project
```

The above is better than the below.

```
github.com/username/PROJECT
or
github.com/username/Project
```

**If the name is long and is a combination of words, you should use an abbreviation or dash, don’t use an underscore.**

Example:

```
github.com/username/postgresusers
or
github.com/username/pgusers
or
github.com/username/postgres-users
```