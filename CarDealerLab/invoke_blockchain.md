# Web App에서 BlockChain 호출하기

## A. Blockchain Query하기

### 애플리케이션에 접속합니다.

1. 앞에서 생성한 Application의 console을 클릭해서 App을 엽니다.
   ![](images/accs7.png)

    이 App은 제조업체인 Detroit와 파트너 업체인 Sam Dealer, Jude Dealer에서 사용하는 Vehicle Trace Dashboard 입니다. 이 App은 각각 Org의 Chaincode에 접근해서 해당 채널에 배포된 Ledger로 부터 Query를 해오고, 다른 업체로 부품을 전달하는 Transaction을 수행할 수 있는 기능을 제공합니다.
  
   딜러들은 거래를 제조업체와 공유하는 비공개 채널로 푸시합니다.  "carTrace"이라는 Smart Contract(체인코드)에 의해 실행됩니다. 이 Smart Contract은 컨소시엄의 허가된 원장에 대한 모든 상태를 관할합니다. 해당 대리점의 "차량 추적 대시 보드"를 열고 포함된 부품 및 차량을 쿼리 할 때 사용하십시오.

### 블럭체인에 접속하기 위한 설정을 합니다.

1. 화면 상단 우측의 아이콘을 클릭합니다.
    ![](images/sample_app_setup1.png)

1. 화면이 뜨면 다음의 내용을 자신의 환경에 맞게 입력합니다.
    - REST Proxy Endpoint : REST API 주소입니다. 이 정보는 Founder Console(DetroitAuto) 에서 확인할 수 있습니다. (예: `https://62C57E2D8944EBB48B1E6C5A28A461.blockchain.ocp.oraclecloud.com:443`)
    - Channel Name : 접속하고자 하는 채널 이름입니다. (예: **samchannel**, **judechannel**)
    - UserName : 접속 아이디 (예: **abc@oracle.com**)
    - Password : 접속 패스워드 (예: `xxxxx`)

    ![](images/sample_app_setup2.png)

1. 블럭체인의 REST Proxy에 접근하기 위한 환경을 완료하였습니다.

### Sam Dealer의 채널로 접근합니다.

1. 화면 상단 우측의 아이콘을 클릭하여 Sam 이 사용하는 채널로 설정합니다.
    - REST Proxy Endpoint : 위에서 설정한 동일 REST URL을 사용합니다.
    - Channel Name : **samchannel**

    ![](images/sample_app_setup3.png)



1. 다음과 같이 제목이 samchannel 로 바뀌었음을 알 수 있습니다.

    ![](images/sample_app_setup4.png)
 

1. 다음 정보를 넣고 쿼리를 해보세요. 

    참고 : 쉼표 ","를 사용하여 쿼리 필드에서 한 번에 여러 항목을 검색 할 수 있습니다.

    에어백 일련 번호:

    | abg1234, abg1235 |
    | -- |

    차량 번호:

    | dtrt10001, dtrt10002 |
    | -- |

    위의 쿼리를 Jude Dealer의 대시 보드에서 수행하면 채널 프라이버시로 인해 결과가 조회되지 않습니다.

### Jude Dealer의 채널로 접근합니다.

1. 화면 상단 우측의 아이콘을 클릭하여 Jude 이 사용하는 채널로 설정합니다.
    - REST Proxy Endpoint : 위에서 설정한 동일 REST URL을 사용합니다.
    - Channel Name : **judechannel**

    ![](images/sample_app_setup5.png)


1. 다음과 같이 제목이 judechannel 로 바뀌었음을 알 수 있습니다.

    ![](images/sample_app_setup6.png)

1. 다음 정보를 넣고 쿼리를 해보세요. 
   
    VehiclePart 일련 번호 : 

    | whl1241, win1242 | 
    | -- |

    > Detroit Auto 샘플 애플리케이션에서는 모든 쿼리를 실행할 수 있습니다.


1. 그리드에서 항목을 선택하면 해당 기록이 표시됩니다. 트랜잭션을 선택하면 항목의 내역 상태가 자세히 표시됩니다.

    > 참고 : 그래프에는 현재 시간보다 45 일 이내에 실행된 트랜잭션만 표시됩니다. 최근 거래를 보려면 그래프의 현재 날짜로 스크롤해야 할 수 있습니다.

    ![](images/sampleapp1.png)

## C. Transaction 실행하기
1. 다시 Sam Dealer 채널로 설정합니다. 화면 상단 우측의 아이콘을 클릭하여 Sam 이 사용하는 채널로 설정합니다.
    - REST Proxy Endpoint : 위에서 설정한 동일 REST URL을 사용합니다.
    - Channel Name : **samchannel**

    ![](images/sample_app_setup3.png)



1. "Vehicle Trace Dashboard"를 열고 **transfer**탭을 선택하여 트랜잭션을 실행하십시오.
    **Sam Dealer**에서 **Detroit Auto**로 차량번호 **dtrt10001**를 전송해 보겠습니다.

1. 사용자 이름 **SamDealer**와 아무 비밀번호나 넣고 로그인하십시오.

    차량 번호로 **dtrt10001**을 입력하십시오(Chassis number of a vehicle). 
    
    Vehicle **new owner**로 **DetroitAuto**를 입력하십시오. 
    
    그런 다음 **Transfer Vehicle** 버튼을 클릭하십시오.

    ![](images/sampleapp_transfer_vehicle.png)

1. 다시 Trace 탭으로 이동해서 **dtrt10001** 를 조회해서 해당 차량의 Transaction History를 확인할 수 있습니다. 
2. 오른쪽의 API Details에서와 같이 API 호출 시에 JSON 형태로 History 데이터가 반환되는 것을 확인할 수 있습니다.
![](images/sample_app_trace.png)

## D. REST API를 통해 Transaction 실행

CarTrace의 모든 트랜잭션과 결과들은 REST API로 호출할 수 있습니다.
REST 호출을 실행할 때 두 가지 옵션이 있습니다.
- Postman과 같은 Tool 사용
- 터미널에서 "curl" 명령 사용

Postman은 앞서 기초 데이터를 넣을 때 이미 사용했습니다.
이미 배포된 곳에 Query를 위한 샘플용 request가 들어 있습니다.

1. 왼쪽 request들 가장 아래 부분에 Query Vehicle 이라는 request를 선택하십시오.
![](images/postman_queryhistory.png)

1. 오른쪽 입력파라미터 중 args로 위에서 사용한 차량번호 **dtrt10001**를 입력합니다.
    send를 누르면 결과가 차량 history에 대한 결과가 json 형태로 반환됩니다.

    위 결과는 curl 명령어를 사용해도 동일한 결과를 얻을 수 있습니다.

---
<b>수고하셨습니다. 모든 Lab 과정을 끝내셨습니다.</b>

[이전 Lab으로 이동](README.md)