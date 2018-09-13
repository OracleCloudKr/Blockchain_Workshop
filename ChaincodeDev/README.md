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

Chaincode는 Hyperleger위에 돌아가는 SmartContract 코드로서 Ledger를 저장할 때 비지니스 로직을 실행하게 할 수 있습니다. Chaincode는 Transaction에서 Ledger를 쿼리하거나 Update하기 위해 호출됩니다.
현재 지원되는 언어는 Go와 Node.js로 작성합니다. (Java 도 지원될 예정)

# Chaincode API란?

모든 Chaincode 프로그램은 Chaincode 인터페이스를 구현해야 합니다.
- [Go](https://godoc.org/github.com/hyperledger/fabric/core/chaincode/shim#Chaincode)
- Node.js

다음은 Chaincode 인터페이스의 주요 메서드들입니다.
- **Init method** : Instantiate 나 upgrade 시에 호출
- **Invoke method** : invoke 트랜잭션 요청을 받을 때 호출

- shim 패키지에는 다른 인터페이스인 chaincodeStubInterface가 있는데 원장에 액세스해서 수정하고 chaincode들 사이에서 호출하는데 사용됩니다.

# Simple Asset Chaincode

간단히 원장에 assets(Key-value 쌍)을 생성하는 기본 chaincode를 만들도록 하겠습니다.

- **코드 저장 위치 선택**

Go 프로그래밍을 하려는 경우는 Go 프로그래밍 언어가 설치되어 있어야 하는데 시스템에 올바르게 설치 되어 있는지 우선 확인해야 합니다.

Go는 1.10.x 이상이 설치되어 있어야 하고 GOPATH가 시스템 환경 변수로 설정 되어 있어야 합니다. 

또한 PATH에 GOPATH/bin 이 잡혀 있어야 합니다.

[Go 다운로드](https://golang.org/dl/)

이제 체인코드 애플리케이션을 $GOPATH/src/에 저장하시기 바랍니다.
해당 디렉토리로 이동하여 sacc 라고 하는 새로운 파일을 하나 생성합니다.

- **Housekeeping**

먼저, 몇 가지 정리 작업부터 시작하겠습니다. 모든 체인 코드와 마찬가지로, 그것은 특히 체인코드 인터페이스, Init 및 Invoke 함수를 구현합니다. 

체인코드에 필요한 의존성에 대한 go import 문을 추가합니다.
체인 코드 shim 패키지와 피어 protobuf 패키지를 가져옵니다. 

다음으로 Chaincode shim 함수의 수신기로 SimpleAsset 구조체를 추가해 보겠습니다
<pre></code>
package main
import (
    "fmt"

    "github.com/hyperledger/fabric/core/chaincode/shim"
    "github.com/hyperledger/fabric/protos/peer"
)
// 저장하기 위한 구조체인 asset을 관리하기 위한 간단한 체인코드 
type SimpleAsset struct {
}
</code></pre>

- **Chaincode 초기화**

다음 Init 함수를 구현할 것입니다.

<pre><code>
// 체인코드가 instantiate될때 호출이 되며, 초기데이터를 설정하기 위해 사용합니다.
func (t *SimpleAsset) Init(stub shim.ChaincodeStubInterface) peer.Response {

}</code></pre>

| 주의 : Chaincode upgrade 역시 이 함수를 호출합니다. 기존 코드를 upgrade할 Chaincode를 작성할 때 Init 함수를 적절하게 수정해야 합니다. 특히 Upgrade의 일부로서 초기화할 것이 아무것도 없는 경우 위와 같이 비어 있는 Init 메소드를 제공해야 합니다. |
| --- |

다음으로 Init 호출될때 특정값들이 정상적으로 들어 왔는지 유효성 검사를 하는 코드를 넣도록 하겠습니다.
입력 받은 값을 얻기 위해서는 ChaincodeStubInterface.GetStringArgs 함수를 사용합니다.
체인코드를 Instantiate 할 때나 Upgrade 할 때 입력하는 값이 이 함수로 전달되게 됩니다.

<pre>
func (t *SimpleAsset) Init(stub shim.ChaincodeStubInterface) peer.Response {
  <b>args := stub.GetStringArgs()</b>
  if len(args) != 2 {
    return shim.Error("Incorrect arguments. Expecting a key and a value")
  }

  // PutState를 호출하여 ledger에 key/value를 저장
  err := <b>stub.PutState</b>(args[0], []byte(args[1]))
  if err != nil {
    return shim.Error(fmt.Sprintf("Failed to create asset: %s", args[0]))
  }
  return shim.Success(nil)
}
</pre>

입력값이 유효하다는 것을(2개의 파라미터 임을) 확인한 후에는 원장(ledger)에 입력 받은 값을 저장합니다. 
이를 위해 인수로 전달 된 Key와 Value값으로 ChaincodeStubInterface.PutState를 호출 할 것입니다. 모든 것이 성공적으로 끝나면 Peer.Response 객체에 Success 값을 넣어 반환합니다.

- **Invoking the Chaincode**

체인코드에 들어 있는 값들을 호출하는 메서드를 구현해보도록 하겠습니다.

<pre><code>
func (t *SimpleAsset) Invoke(stub shim.ChaincodeStubInterface) peer.Response {
}</code></pre>

Invoke 함수의 파라미터는 [function이름+파라미터 배열] 로 구성되어 있습니다.

이 값들은 ChaincodeStubInterface.GetFunctionAndParameters를 호출하여 함수 이름과 파라미터를 추출할 수 있습니다.

이 예제에서는 function을 get과 set 두 가지 로만 제한하도록 하겠습니다.

<pre><code>
func (t *SimpleAsset) Invoke(stub shim.ChaincodeStubInterface) peer.Response {
    <b>fn, args := stub.GetFunctionAndParameters()</b>

}</code></pre>

fn 변수에 함수이름이 들어가게 되고, args에는 나머지 파라미터들이 들어가게 됩니다.
fn의 함수명에 맞게 set과 get 함수들을 별도로 호출하도록 구현해보겠습니다.

<pre><code>
func (t *SimpleAsset) Invoke(stub shim.ChaincodeStubInterface) peer.Response {
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

    // 각 함수에서 받은 결과를 Success payload로 실어서 반환
    return shim.Success([]byte(result))
}
</code></pre>

- **Chaincode 애플리케이션 구현**

이제 get, set 함수들을 구현해 보겠습니다. 

원장의 상태에 액세스하려면 chaincode shim API의 ChaincodeStubInterface.PutState 및 ChaincodeStubInterface.GetState 함수를 활용해야 합니다.

<pre><code>
// ledger에 key/value를 저장합니다. 만약 키가 있으면 새로운 값으로 override 합니다.
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

// 특정 key에 대한 value를 반환합니다.
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

- **메인함수 추가**

마지막으로 shim.Start 함수를 호출 할 main 함수를 추가해야 합니다.

<pre><code>
package main

import (
    "fmt"

    "github.com/hyperledger/fabric/core/chaincode/shim"
    "github.com/hyperledger/fabric/protos/peer"
)

type SimpleAsset struct {
}

...

func main() {
    if err := shim.Start(new(SimpleAsset)); err != nil {
            fmt.Printf("Error starting SimpleAsset chaincode: %s", err)
    }
}
</code></pre>

- **메인함수 추가**
[체인 코드 프로그램 소스 링크](./sacc)


- **Chaincode 빌드**


- **Chaincode 실행**
이제 a 값이 20으로 변경하기 위한 호출을 실행합니다.

| peer chaincode invoke -n mycc -c &#39;{&quot;Args&quot;:[&quot;set&quot;, &quot;a&quot;, &quot;20&quot;]}&#39; -C myc |
| --- |

마지막으로 a가 20으로 변경되었는지 쿼리해봅니다.

| peer chaincode query -n mycc -c &#39;{&quot;Args&quot;:[&quot;query&quot;,&quot;a&quot;]}&#39; -C myc |
| --- |