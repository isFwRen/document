### Ethereum Development whit go

#### Introduction

**用go来做以太坊开发**

这本迷你书的本意是给任何想用go进行以太坊开发的同学一个概括的介绍。本意是如果你已经对以太坊和Go有一些熟悉，但是对于怎么把两者结合起来还有些无从下手，那么这本书就是一个好的起点。你会学习然后使用go与智能合约交互，还有如何完成一些日常的查询任务。

**在线电子书**

[https://goethereumbook.org](https://goethereumbook.org/)

**介绍**

> 以太坊是一个开源、公开，基于区块链的分布式计算平台和具备智能合约(脚本)功能的操作系统。它通过基于交易的状态转移支持中本聪共识的一个改进算法。

**维基百科**

以太坊是一个区块链，允许开发者创建完全去中心化运行的应用程序，这意味着没有单个实体可以将其删除或修改它。部署到以太坊上的每个应用都用以太坊网络上每个完整客户端执行。

**Solidity**

Solidity是一种用于编写智能合约的图灵完备变成语言。Solidity被编译成以太坊虚拟机可执行的字节码。

**go-ethereum**

本书中，我们将使用Go官方以太坊实现go-ethereum来和以太坊区块链进行交互。go-ethereum，也被简称**Geth**，是最流行的以太坊客户端。因为它是用Go开发的，当使用Golang开发应用程序时，Geth提供了读写区块的一切功能。

**Block Explorers**

Eherscan是一个用于查看和深入研究区块链上数据的网站。这些类型的网站被称为区块浏览器，因为它们允许你查看区块(包含交易)的内容。区块是区块链的基础构成要素。区块包含在已分配的出块时间内开采的所有交易数据。区块浏览器也运行你查看智能合约执行期间释放的事件以及诸如支付的gas和交易的以太币数量等。

我们还讲深入研究蜂群（Swarm）和耳语（Whisper），分别是一个文件存储协议和一个点对点的消息传递协议，它们是实现完全去中心化和分布式应用程序需要的另外两个核心。

![image-20240607223716201](https://fwr-typora-img.oss-cn-guangzhou.aliyuncs.com/typroa-img/image-20240607223716201.png)

#### 客户端

客户端是以太坊网络的入口。客户端需要广播交易和读取区块链数据

##### 创建客户端

**初始化客户端**

用go初始化以太坊客户端是和区块链交互所需的基本步骤。首先，导入go-etherem的ethclient包，并通过调用区块链服务提供者URL的**`Dial`**来初始化它。

若你没有现有的以太坊客户端，你可以连接到infura网关。Infura管理着一批安全，可靠，可拓展的以太坊[geth和parity]节点，并且在接入以太坊网络时降低了新人的入门门槛。

两种方式：

```
client, err := ethclient.Dial("https://cloudflare-eth.com")
```

你还可以将路径传递给IPC端点文件。

```
client, err := ethclient.Dial("/home/user/.ethereum/geth.ipc")
```

`ethclient.Dial` 函数通过传入 **IPC 文件路径**来建立与本地 Geth 节点的连接。**IPC（Inter-Process Communication）文件**允许本地进程之间的通信，这在以太坊中用于客户端与节点之间的通信。

**与 Geth 节点交互**：

通过建立 IPC 连接，你的 Go 应用程序可以与 Geth 节点交互，发送 JSON-RPC 请求，进行操作如查询区块、发送交易、调用智能合约等。

**高效通信**：

- IPC 通道在本地环境中比 HTTP 更加高效，因为它使用了 Unix 套接字进行进程间通信，适合高频率的 RPC 调用和数据交互。

**`ethclient.Dial` 用法示例**

下面是一个完整的 Go 代码示例，展示如何使用 `ethclient.Dial` 连接到本地 Geth 节点并获取最新区块号：

```go
package main

import (
    "context"
    "fmt"
    "log"
    "github.com/ethereum/go-ethereum/rpc"
    "github.com/ethereum/go-ethereum/ethclient"
)

func main() {
    // 连接到本地 Geth 节点的 IPC 文件
    client, err := ethclient.Dial("/home/user/.ethereum/geth.ipc")
    if err != nil {
        log.Fatalf("Failed to connect to the Ethereum client: %v", err)
    }

    // 获取最新的区块号
    blockNumber, err := client.BlockNumber(context.Background())
    if err != nil {
        log.Fatalf("Failed to retrieve the latest block number: %v", err)
    }

    fmt.Printf("Latest block number: %d\n", blockNumber)
}

```

**使用Ganache**

Ganache（正式名称为testrpc）是一个用Node.js编写的以太坊实现，用于在本地开发者去中心化应用程序进行测试。现在我们将带着你完成安装并连接到它。

首先通过NPM安装ganache。

```
npm install -g ganache-cli
```

然后运行ganache cli客户端。

```
ganache-cli
```

现在连接到`http://localhost:8584`上的ganache RPC主机

```
client, err := ethclient.Dial("http://localhost:8545")
if err != nil {
  log.Fatal(err)
}
```

在启动ganache时，你还可以使用相同的助记词来生成相同序列的公开地址。

```
ganache-cli -m "much repair shock carbon improve miss forget sock include bullet interest solution"
```

我强烈推荐您通过阅读其[文档](http://truffleframework.com/ganache/)熟悉ganache。

****

**完整代码**

`client.go`

```go
package main

import (
    "fmt"
    "log"

    "github.com/ethereum/go-ethereum/ethclient"
)

func main() {
    client, err := ethclient.Dial("https://cloudflare-eth.com")
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println("we have a connection")
    _ = client // we'll use this in the upcoming sections
}
```

#### 账户

以太坊上的账户要么是**钱包地址**，要么是**智能合约地址**。它们看起来像。是你用来向另一个用户发送ETH的东西，也用于在需要与区块链交互时引用区块链上的智能合约。它们是唯一的，并且派生自私钥。我们将在后面的部分中更深入的介绍私钥/公钥对。

`0x71c7656ec7ab88b098defb751b7401b5f6d8976f`

为了在go-ethereum中使用账户地址，你必须首相将它们转换为go-ethereum类型。

`common.Address`

```go
address := common.HexToAddress("0xb199B02A8eE9610e257574853857a0fD577fd67A")
fmt.Println(address.Hex())
```

几乎你会在任何地方使用这种类型，你可以将以太坊地址传递给来自go-ethereum的方法。现在你已经了解了账户和地址的基础知识，让我们在下一节中学习如何检索ETH账户余额。

****

完整代码

address.go

```go
package main

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
)

func main() {
	address := common.HexToAddress("0xb199B02A8eE9610e257574853857a0fD577fd67A")
	fmt.Println("address:", address.Hex())
	fmt.Println("address:", address.Bytes())
}

```

##### 账户余额

读取账户的余额非常简单，调用客户端的方法，向其传递账户地址和可选块号。设置为块号将返回最新的余额。`BlalanceAt nil`

通过区块编号，你可以读取该区块识的账户余额。块号必须是`big.Int`

```go
blockNumber := big.NewInt(5532993)
if err != nil {
    log.Fatal(err)
}

fmt.Println(balance) //25893180161173005034
```

以太坊中的数字使用尽可能小的单位进行处理，因为它们的定点精度，在ETH的情况下是wei。要读取ETH值，你必须进行计算。因为我们要处理大数字，所以我们必须导入原生Go和包。这是你如何进行转换的。`wei/10^18 math math/big`

```
pendingBalance,err := client.PrendingBalanceAt(context.Background(),account)
fmt.Println(pendingBalance)
```



****

完整代码

account_balance.go

```go
package main

import (
	"context"
	"fmt"
	"log"
	"math"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	client, err := ethclient.Dial("https://cloudflare-eth.com")
	if err != nil {
		log.Fatal(err)
	}

	account := common.HexToAddress("0x71c7656ec7ab88b098defb751b7401b5f6d8976f")
	balance, err := client.BalanceAt(context.Background(), account, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(balance) // 25893180161173005034

	blockNumber := big.NewInt(5532993)
	balanceAt, err := client.BalanceAt(context.Background(), account, blockNumber)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(balanceAt) // 25729324269165216042

	fbalance := new(big.Float)
	fbalance.SetString(balanceAt.String())
	ethValue := new(big.Float).Quo(fbalance, big.NewFloat(math.Pow10(18)))
	fmt.Println(ethValue) // 25.729324269165216041

	pendingBalance, err := client.PendingBalanceAt(context.Background(), account)
	fmt.Println(pendingBalance) // 25729324269165216042
}

```

##### 账户代币余额

要了解如何读取账户[代币](https://goethereumbook.org/GLOSSARY.html#token) （ERC20） 余额，请前往 [ERC20 代币智能合约部分](https://goethereumbook.org/smart-contract-read-erc20)。
