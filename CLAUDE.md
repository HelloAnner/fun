# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

这是一个基于Go语言的算法和数据结构仓库，使用leetgo工具管理LeetCode题目。仓库包含：

- **LeetCode Solutions**: 在 `/go/` 目录中按题目编号组织的解决方案
- **Data Structures**: 核心数据结构定义在 `/data_structures/`
- **Engineering**: 专门的实现如布隆过滤器在 `/engineering/`
- **Leetgo Configuration**: 通过 `leetgo.yaml` 配置文件管理

## Development Commands

**Leetgo Commands:**
```bash
leetgo pick <problem-id>    # 获取并生成LeetCode题目
leetgo test                 # 运行当前题目的测试
leetgo submit               # 提交解决方案
```

**Go Commands:**
```bash
go build ./...              # 构建所有包
go fmt ./...               # 格式化所有Go代码
go mod tidy                # 清理模块依赖
```

**Individual Problem Testing:**
```bash
cd go/<problem-folder>      # 进入具体题目目录
go run solution.go          # 运行解决方案
```

## Code Structure

**Main Directories:**
- `go/` - LeetCode题目解决方案，按题目编号组织（如 `0001.two-sum/`）
- `data_structures/` - 核心数据结构定义
- `engineering/` - 专门的工程实现

**Key Data Structures:**
- `ListNode` - 单链表节点 (`data_structures/list_node.go`)
- `TreeNode` - 二叉树节点 (`data_structures/tree_node.go`)
- `GraphNode` - 图节点 (`data_structures/graph_node.go`)
- `BitArray` - 位数组 (`data_structures/bit_array.go`)

**LeetCode Solution Structure:**
每个题目目录包含：
- `solution.go` - 主要解决方案文件
- `question.md` - 题目描述（中文）
- `testcases.txt` - 测试用例

## Leetgo Configuration

项目使用 `leetgo.yaml` 配置：
- 默认语言：Go
- 题目语言：中文 (zh)
- 输出目录：`go/`
- LeetCode网站：leetcode.cn

## Code Style

- Go 1.22语法和约定
- 中文注释和题目描述
- 每个题目独立的包结构（package main）
- 使用leetgo提供的测试工具
- 遵循LeetCode题目命名约定