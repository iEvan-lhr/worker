<p align="center">
    <h1 align="center">Worker</h1>
</p>


## 介绍

使用Golang构建的工作流，任务调度系统使用<a href="https://github.com/iEvan-lhr/nihility-dust">nihility-dust</a>.


## Introduce

Worker使用了全新的协程调度器，将所有的工作流交给协程来做，同时支持多任务的中断控制，也能发挥CPU的全部核心性能，而并非单核性能。

## 快速开始


```bash
go get github.com/iEvan-lhr/worker
```


## 关于 Worker

从一个天真想法的全协程执行控制器开始，到Worker工作流调度器的开发，一步一个脚印，也欢迎所有人对Worker进行改造和使用.

