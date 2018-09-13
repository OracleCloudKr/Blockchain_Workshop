# ChainCode 개발하기 #
-----
## 사전 준비사항 ##
+ Oracle Cloud Trial Account ([Trial 계정 신청하기](http://www.oracloud.kr/post/oracle_cloud_trial_universal/))
+ 이전 Lab인 CarDealer Lab을 완료하거나, BlockChain Cloud Service에서 Provisioning한 인스턴스 필요
+ 개발 툴 ([Visual Studio Code](https://code.visualstudio.com/))

## 소개 ##
이번 Lab에서는 체인코드를 개발하는 방법에 대해 학습하고 간단한 체인코드를 개발해서 Oracle Blockchain Cloud Service에 배포하는 과정을 실습합니다.

# Lab 시작

## 1. 

# Hyperledger Chaincode 시작하기























참고 문헌 : http://hyperledger-fabric.readthedocs.io/en/latest/index.html

# Chaincode란?

Chaincode는 Go와 Node.js로 작성된 프로그램이다.(Server 및 Client 프로그램으로 나뉘어 있다. Java 도 지원될 예정임)
  - Go
  - Node.js

# Chaincode API란?

모든 Chaincode 프로그램은 Chaincode 인터페이스를 구현해야 합니다.

메소드는 수신된 트랜잭션에 응답하여 호출되며 다음은 Chaincode에 대한 간략한 설명입니다.

 - **Init method** : Instantiate 나 upgrade 트랜잭션 요청을 받을 때 호출
  - **Invoke method** : invoke 트랜잭션 요청을 받을 때 호출
  - 다른 chaincode &quot;shim&quot; API들 안에 다른 인터페이스는 chaincodeStubInterface가 있는데 원장에 액세스해서 수정하고 chaincode들 사이에서 호출하는데 사용됩니다.
  - 이 문서를 통하여 간단한 &quot;Assets&quot;을 관리하는 간단한 체인코드 애플리케이션을 구현하여 이러한 API 사용법을 보여줍니다.


# Simple Asset Chaincode

이 애플리케이션은 원장에 assets(Key-value pairs)을 생성하는 기본 샘플 chaincode입니다.

- **코드 저장 위치 선택**

Go 프로그래밍을 하려는 경우는 Go 프로그래밍 언어가 설치되어 있어야 하는데 시스템에 올바르게 설치 되어 있는지 우선 확인해야 합니다.

지금, 체인코드 애플리케이션을 $GOPATH/src/에 저장할 것입니다.

쉽게 하기 위해 보통 다음과 같은 커맨드를 사용합니다.

| mkdir -p $GOPATH/src/sacc &amp;&amp; cd $GOPATH/src/sacc |
| --- |

소스 작성을 위해 다음과 같이 파일을 생성합니다.

| touch sacc.go |
| --- |

- **Housekeeping**

먼저, 몇 가지 정리 작업부터 시작하겠습니다. 모든 체인 코드와 마찬가지로, 그것은 특히 체인코드 인터페이스, Init 및 Invoke 함수를 구현합니다. 체인코드에 필요한 의존성에 대한 go import 문을 추가해 봅시다. 체인 코드 심 패키지와 피어 protobuf 패키지를 가져옵니다. 다음으로 Chaincode shim 함수의 수신기로 SimpleAsset 구조체를 추가해 보겠습니다

<pre><code>package main
import (
    "fmt"

    "github.com/hyperledger/fabric/core/chaincode/shim"
    "github.com/hyperledger/fabric/protos/peer"
)
// SimpleAsset implements a simple chaincode to manage an asset
type SimpleAsset struct {
}
</code></pre>

- **Chaincode 초기화**

다음 우리는 Init 함수를 구현할 것입니다.

<pre><code>
// Init is called during chaincode instantiation to initialize any data.
func (t *SimpleAsset) Init(stub shim.ChaincodeStubInterface) peer.Response {

}</code></pre>

% 주의 : Chaincode upgrade 역시 이 함수를 호출합니다. 기존 코드를 upgrade할 Chaincode를 작성할 때 Init 함수를 적절하게 수정해야 합니다. 특히 Upgrade의 일부로서 초기화할 것이 아무것도 없는 경우 위와 같이 비어 있는 Init 메소드를 제공해야 합니다.



다음은 ChaincodeStubInterface.GetStringArgs 함수를 사용하여 Init 호출에 대한 인수를 검색하고 유효성을 검사합니다. 여기서는 key-value pair 기대합니다.

<pre>
// Init is called during chaincode instantiation to initialize any
// data. Note that chaincode upgrade also calls this function to reset
// or to migrate data, so be careful to avoid a scenario where you
// inadvertently clobber your ledger's data!
func (t *SimpleAsset) Init(stub shim.ChaincodeStubInterface) peer.Response {
  // Get the args from the transaction proposal
  args := stub.GetStringArgs()
  if len(args) != 2 {
    return shim.Error("Incorrect arguments. Expecting a key and a value")
  }

  // Set up any variables or assets here by calling stub.PutState()

  // We store the key and the value on the ledger
  err := stub.PutState(args[0], []byte(args[1]))
  if err != nil {
    return shim.Error(fmt.Sprintf("Failed to create asset: %s", args[0]))
  }
  return shim.Success(nil)
}
</pre>
이제 호출이 유효하다는 것을 확인 했으므로 원장에 초기 상태를 저장합니다. 이를 위해 우리는 인수로 전달 된 Key와 Value값으로 ChaincodeStubInterface.PutState를 호출 할 것입니다. 모든 것이 잘되었다고 가정하면 Peer를 반환합니다.

- **Invoking the Chaincode**

먼저 Invoke 함수의 서명을 추가합니다.

<pre><code>
// Invoke is called per transaction on the chaincode. Each transaction is
// either a 'get' or a 'set' on the asset created by Init function. The 'set'
// method may create a new asset by specifying a new key-value pair.
func (t *SimpleAsset) Invoke(stub shim.ChaincodeStubInterface) peer.Response {

}</code></pre>

위의 Init 함수 에서처럼, 우리는 ChaincodeStubInterface에서 인수를 추출 해야합니다. Invoke 함수의 인수는 호출 할 체인코드 애플리케이션 함수의 이름입니다. 이경우 애플리케이션은 단순히 자산의 가치를 설정하거나 현재 상태를 검색 할 수 있는 두 가지 기능, set과 get을 가질 것입니다.

먼저 ChaincodeStubInterface.GetFunctionAndParameters를 호출하여 함수 이름과 해당 chaincode 응용 프로그램 함수에 대한 매개 변수를 추출합니다.

<pre><code>
// Invoke is called per transaction on the chaincode. Each transaction is
// either a 'get' or a 'set' on the asset created by Init function. The Set
// method may create a new asset by specifying a new key-value pair.
func (t *SimpleAsset) Invoke(stub shim.ChaincodeStubInterface) peer.Response {
    // Extract the function and args from the transaction proposal
    fn, args := stub.GetFunctionAndParameters()

}</code></pre>

그런 다음 함수 이름을 set 또는 get 중 하나로 확인하고 해당 chaincode 애플리케이션 함수를 호출하여 gRPC 프로토콜 메시지로 응답을 serialize하는 shim.Success 또는 shim.Error 함수를 통해 적절한 응답을 반환합니다.

<pre><code>
// Invoke is called per transaction on the chaincode. Each transaction is
// either a 'get' or a 'set' on the asset created by Init function. The Set
// method may create a new asset by specifying a new key-value pair.
func (t *SimpleAsset) Invoke(stub shim.ChaincodeStubInterface) peer.Response {
    // Extract the function and args from the transaction proposal
    fn, args := stub.GetFunctionAndParameters()

    var result string
    var err error
    if fn == "set" {
            result, err = set(stub, args)
    } else {
            result, err = get(stub, args)
    }
    if err != nil {
            return shim.Error(err.Error())
    }

    // Return the result as success payload
    return shim.Success([]byte(result))
}
</code></pre>

- **Chaincode 애플리케이션 구현**

언급했듯이, 우리의 chaincode 애플리케이션은 Invoke 함수를 통해 호출 할 수 있는 두 가지 기능을 구현합니다. 이제 이러한 함수들을 구현해 보겠습니다. 위에서 언급했듯이 원장의 상태에 액세스하려면 chaincode shim API의 ChaincodeStubInterface.PutState 및 ChaincodeStubInterface.GetState 함수를 활용하십시오.

<pre><code>
// Set stores the asset (both key and value) on the ledger. If the key exists,
// it will override the value with the new one
func set(stub shim.ChaincodeStubInterface, args []string) (string, error) {
    if len(args) != 2 {
            return "", fmt.Errorf("Incorrect arguments. Expecting a key and a value")
    }

    err := stub.PutState(args[0], []byte(args[1]))
    if err != nil {
            return "", fmt.Errorf("Failed to set asset: %s", args[0])
    }
    return args[1], nil
}

// Get returns the value of the specified asset key
func get(stub shim.ChaincodeStubInterface, args []string) (string, error) {
    if len(args) != 1 {
            return "", fmt.Errorf("Incorrect arguments. Expecting a key")
    }

    value, err := stub.GetState(args[0])
    if err != nil {
            return "", fmt.Errorf("Failed to get asset: %s with error: %s", args[0], err)
    }
    if value == nil {
            return "", fmt.Errorf("Asset not found: %s", args[0])
    }
    return string(value), nil
}
</code></pre>

- **전체 코드**

마지막으로 shim.Start 함수를 호출 할 main 함수를 추가해야 합니다. 다음은 전체 체인 코드 프로그램 소스입니다.

<pre><code>
package main

import (
    "fmt"

    "github.com/hyperledger/fabric/core/chaincode/shim"
    "github.com/hyperledger/fabric/protos/peer"
)

// SimpleAsset implements a simple chaincode to manage an asset
type SimpleAsset struct {
}

// Init is called during chaincode instantiation to initialize any
// data. Note that chaincode upgrade also calls this function to reset
// or to migrate data.
func (t *SimpleAsset) Init(stub shim.ChaincodeStubInterface) peer.Response {
    // Get the args from the transaction proposal
    args := stub.GetStringArgs()
    if len(args) != 2 {
            return shim.Error("Incorrect arguments. Expecting a key and a value")
    }

    // Set up any variables or assets here by calling stub.PutState()

    // We store the key and the value on the ledger
    err := stub.PutState(args[0], []byte(args[1]))
    if err != nil {
            return shim.Error(fmt.Sprintf("Failed to create asset: %s", args[0]))
    }
    return shim.Success(nil)
}

// Invoke is called per transaction on the chaincode. Each transaction is
// either a 'get' or a 'set' on the asset created by Init function. The Set
// method may create a new asset by specifying a new key-value pair.
func (t *SimpleAsset) Invoke(stub shim.ChaincodeStubInterface) peer.Response {
    // Extract the function and args from the transaction proposal
    fn, args := stub.GetFunctionAndParameters()

    var result string
    var err error
    if fn == "set" {
            result, err = set(stub, args)
    } else { // assume 'get' even if fn is nil
            result, err = get(stub, args)
    }
    if err != nil {
            return shim.Error(err.Error())
    }

    // Return the result as success payload
    return shim.Success([]byte(result))
}

// Set stores the asset (both key and value) on the ledger. If the key exists,
// it will override the value with the new one
func set(stub shim.ChaincodeStubInterface, args []string) (string, error) {
    if len(args) != 2 {
            return "", fmt.Errorf("Incorrect arguments. Expecting a key and a value")
    }

    err := stub.PutState(args[0], []byte(args[1]))
    if err != nil {
            return "", fmt.Errorf("Failed to set asset: %s", args[0])
    }
    return args[1], nil
}

// Get returns the value of the specified asset key
func get(stub shim.ChaincodeStubInterface, args []string) (string, error) {
    if len(args) != 1 {
            return "", fmt.Errorf("Incorrect arguments. Expecting a key")
    }

    value, err := stub.GetState(args[0])
    if err != nil {
            return "", fmt.Errorf("Failed to get asset: %s with error: %s", args[0], err)
    }
    if value == nil {
            return "", fmt.Errorf("Asset not found: %s", args[0])
    }
    return string(value), nil
}

// main function starts up the chaincode in the container during instantiate
func main() {
    if err := shim.Start(new(SimpleAsset)); err != nil {
            fmt.Printf("Error starting SimpleAsset chaincode: %s", err)
    }
}
</code></pre>



- **Chaincode 빌드**

이제 a 값이 20으로 변경하기 위한 호출을 실행합니다.

| peer chaincode invoke -n mycc -c &#39;{&quot;Args&quot;:[&quot;set&quot;, &quot;a&quot;, &quot;20&quot;]}&#39; -C myc |
| --- |

마지막으로 a가 20으로 변경되었는지 쿼리해봅니다.

| peer chaincode query -n mycc -c &#39;{&quot;Args&quot;:[&quot;query&quot;,&quot;a&quot;]}&#39; -C myc |
| --- |