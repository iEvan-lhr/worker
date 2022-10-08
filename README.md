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

## 简单模式
<h3>使用</h3>

```bash
anything.AddEasyMission()
```
<h3>来注册方法</h3>

<h3>支持</h3>
func()<br />
Struct<br />
<h3>两种模式注册方法</h3>

<h3>使用</h3>

```bash
    temp := <-anything.DoChanN("方法名称","方法参数")
```
<h3>来执行方法</h3>
<h3>temp会接收返回的参数 为any类型</h3>
<h3>通过Pursuit下标的方式取出返回值</h3>

```bash
    db := temp.Pursuit[0].(*gorm.DB)
```
<h2>优点</h2>
<h3>高度解耦</h3>
<h3>避免循环依赖</h3>
<h2>缺点</h2>
<h3>执行速度比正常执行慢约0.5个数量级</h3>

## 进阶模式
<h3>使用</h3>

```bash
    e := engine.Engine{
		W: anything.Wind{},
	}
    e.Start("端口", []any{"方法注册:结构体"}, []any{"路由注册:结构体"})
```
<h3>来开启路由全代理模式</h3>
<h2>优点</h2>
<h3>高度解耦</h3>
<h3>执行速度高于顺序执行模式</h3>
<h3>避免循环依赖</h3>
<h2>缺点</h2>
<h3>代码可读性不强</h3>
<h3>对代码入参和返回参数有较强限制</h3>

## 详细文档
暂无
## 关于 Worker

从一个天真想法的全协程执行控制器开始，到Worker工作流调度器的开发，一步一个脚印，也欢迎所有人对Worker进行改造和使用.

