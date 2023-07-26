---
layout: default
title: Integration
nav_order: 2
---

# Integration
{: .no_toc }

philo are modular project, which means you can integrate it with other tools.
{: .fs-6 .fw-300 }

<details open markdown="block">
  <summary>
    Table of contents
  </summary>
  {: .text-delta }
- TOC
{:toc}
</details>

---

philo are modular project, which means you can integrate it with other tools. For example, to enable centralized grading, 
you can integrate it with [shopiea](https://github.com/Kyuubang/shopiea) project. This page will explain how to integrate
philo with other tools.

## Integration with shopiea

shopiea is a simple api server that can be used to manage assignment and grading. shopiea is designed to work with philo,
which means you can use shopiea to manage assignment and grading, and use philo to automatically check your assignment.

philo has built-in integration with shopiea, you can use `philo admin` command to manage assignment and grading. for more
information about `philo admin` command, you can read [this](/docs/admin).

build philo to enable admin command

```bash
go build -o philo -tags admin
```
