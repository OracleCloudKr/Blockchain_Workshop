# Car Dealer Demo #
-----
## 사전 준비사항 ##
+ Oracle Cloud Trial Account

## 소개 ##
이번 Lab은 자동차 Oracle BlockChain를 활용하여 Car Dealer 회사에서 자동차 매매에 투명성을 확보할 수 있도록 해주는 Lab 입니다.

## Lab 시나리오
대형 자동차 회사는 딜러 / 구매자의 공급망을 관리하는 소프트웨어(SCM)을 원장 (blockchain)에 권한을 부여하는 방식으로 공급망 시스템을 간소화하는 것을 목표로 하고 있습니다. 최종 목표는 조정하는 노력과 차량 및 부품에 대한 감사 추적 설정과 관련된 전반적인 마찰을 줄이는 것입니다.
딜러가 서로 경쟁 할 수 있으므로 분명히 다른 딜러가 거래의 세부 정보를 보기를 원하지 않습니다 (또는 심지어 거래를 하고 있음에도 불구하고). Blockchain 클라우드 서비스에서는 채널을 사용하여 트랜잭션을 서로 분리하고 해당 채널을 볼 수 있는 동료(peer)만 해당 트랜잭션을 볼 수 있게 함으로써 이를 해결해줍니다.
시나리오에서는 아래 그림과 같이 중앙에 제조업체 (Detroit Auto)와 두 명의 딜러 (Sam 과 Jude를 만들고 채널별로 
구분합니다.

![](images/scenario.png)

이 채널들은 차량 제조업체와 특정 딜러 만 볼 수 있는 별도의 원장을 각각 갖게 됩니다. 채널은 2 개가 아닌 여러 참여자를 가질 수 있고, 네트워크는 더 많은 조직 (예 : 여러 다른 부품 제조업체)을 가질 수 있지만, 이 시나리오에서는 간단한 모델만으로도 충분합니다.

# Lab 시작

## 1. Multi-party Blockchain Network 구성하기
### 인스턴스 생성하기

1. OABCS(Oracle BlockChain Cloud Service) Console 에 접속하기  
OABCS에 접속하기 위해서는 Cloud Console에서 다음과 같이 선택해서 Open Service Console을 클릭합니다.
![](images/gotoabcs.png)

2. Organization(조직) 생성하기 Provisioning
처음 접속을 한 경우에는 아래와 같이 나오게 됩니다.  
오른쪽에 있는 인스턴스 생성 버튼을 누릅니다.
![](images/abcs_main.png)

3. 이제 메인이 되는 DetroitAuto의 BlockChain Network를 만들 예정입니다.  
서비스 이름을 detroitauto 라고 입력한 후 Next 버튼을 누릅니다.  
<span style="color:red">서비스 이름은 모두 소문자로 입력하셔야 합니다.</span>
![](images/create_instance1.png)

4. Next를 누른후 다음 화면에서 Confirm 버튼을 누릅니다.
몇 분 후에 아래와 같이 인스턴스가 생성이 되게 됩니다.
![](images/create_instance2.png)
5. 오른쪽의 햄버거 버튼을 클릭 한 후에 Autonomous Blockchain Console을 선택합니다.

6. 새롭게 생성된 detroitauto Founder의 Console 화면으로 이동합니다.
![](images/detroit_main.png)

7. 다시 Autonomous Blockchain Cloud 화면으로 이동합니다.
![](images/create_participant.png)

8. 이번에는 다른 Dealer들을 Participant(Organization)로 만드는 과정을 구성해보도록 하겠습니다. 
인스턴스 생성을 클릭해서 생성화면으로 이동한 후 아래와 같이 정보를 입력합니다.
먼저 Jude Dealer를 생성하도록 하겠습니다.
여기서는 새로운 네트워크를 생성하는 Founder가 아니기 때문에 Create a new Network 체크를 해제 합니다.
![](images/create_judedealer.png)

9. 동일한 과정을 거쳐 Sam Dealer 를 생성하도록 하겠습니다.
![](images/create_samdealer.png)


#### Founder에 JudeDealer Participant 조직(Organization)을 Join하기
1. 네트워크가 생성되면 각각의 관리 콘솔에 액세스 할 수 있습니다.  
먼저 judedealer 의 콘솔에 접속합니다.
![](images/goto_judedealer.png)

2. Founder는 Participant Org의 두 가지 정보가 필요합니다. 첫째는 통신의 유효성을 확인할 수 있는 인증서 이고, 두번째는 Org에서 관리하는 피어 노드에 대한 정보입니다. 피어 정보는 이후 프로세스까지 필요하지 않습니다.
인증서를 얻으려면 Participant Org의 콘솔을 엽니다. 네트워크 탭에서 햄버거 메뉴를 클릭하고 딜러 조직에서 '인증서 내보내기'를 선택하여 인증서를 내보낼 수 있습니다. 내보낸 파일을 저장하십시오.
![](images/judedealer_main.png)

3. Export 버튼을 눌러 judedealer-certificates.json 파일을 서버로 부터 로컬 컴퓨터에 다운로드 합니다.  
접속하고자 하는 Founder의 Orderer 정보를 import 해야 합니다. Import Orderer Settings를 클릭합니다.
![](images/judedealer_join1.png)

4. "DetroitAuto" Founder 조직의 콘솔로 이동해서 "네트워크" 탭을 선택하십시오. 그런 다음 "DetroitAuto" 조직의 햄버거 메뉴를 클릭하여 Orderer 설정파일을(detroitauto-orderer-settings.json) 내보냅니다. 내보낸 파일을 저장하십시오.
![](images/founder_exportorderer.png)

5. 다시 judedealer 조직의 콘솔로 이동해서 Upload Orderer Settings 를 클릭합니다.
![](images/judedealer_join2.png)
6. 위에서 다운로드 한 detroitauto-orderer-settings.json 를 선택합니다.
![](images/judedealer_select_ordererfile.png)

7. 선택한 파일이 정상적으로 upload 되면, 이 파일을 Import를 클릭해서 반영합니다. 마지막으로 Complete 단계를 클릭합니다.
![](images/judedealer_join4.png)

8. 이제 Judedealer의 메인화면이 정상적으로 열립니다.
![](images/judedealermain.png)

9. 이제 지금까지 JudeDealer를 구성했고, JudeDealer에 대한 정보도 얻었습니다. 일반적으로 이 프로세스는 단일 사용자가 수행하지 않으며 대신 새로운 조직을 Blockchain 네트워크에 탑재하기 위해 메타 데이터를 전송해야 합니다. Oracle Blockchain Cloud Service는 서로 알려진 조직에게만 액세스 권한이 부여되는 'permissioned' 네트워크이며, 지금 하고 있는 작업을 통해 활성화가 되게 됩니다.

    Founder에도 JudeDealer를 추가해주어야 합니다. Founder 콘솔의 Network로 이동해서 Add Organizations 버튼을 클릭합니다.
![](images/founder_addorg.png)

10. Upload Organization Certificates 버튼을 클릭한 후 앞에서 다운로드 한 judedealer-certificates.json 파일을 선택하고 Add 버튼을 눌러 Org를 추가합니다.
![](images/founder_addorg2.png)

11. 다음과 같이 judedealer 가 Participant로 정상적으로 등록된 것을 확인할 수 있습니다.
![](images/founder_addorg3.png)

### Founder에 SamDealer Participant 조직(Organization)을 Join하기
위와 동일한 방식으로 SamDealer Org를 Founder에 추가합니다. 

### Channel
Multi-Party 블록 체인 네트워크를 설정하는 마지막 단계는 각 딜러의 거래를 다른 딜러의 거래와 분리하기 위해 각 딜러의 채널을 만드는 것입니다. 이 시나리오에서는 두 명의 각 대리점을 위한 별도의 채널을 만들 것입니다.

1. Channel 만들기  
이를 생성하려면 Founder Org의 콘솔에서 "Channels" 탭으로 이동하십시오. 그런 다음 "Create a New Channel"를 클릭하십시오. 이 딜러는 '<dealer> channel'이라는 채널 (예 : 'samchannel'과 'judechannel')에 대한 채널을 만들면 됩니다. 각 채널에 대해 ReaderWriter가 포함되어 있고 Detroit Auto Peer의 피어가 모두 채널에 가입되어 있는지 확인하십시오.
![](images/create_newchannel.png)

2. Participant nodes에서 채널에 참가하기  
이 채널들은 현재 Detroit Auto의 피어 노드만을 포함합니다. 딜러는 참여하는 노드가 없으므로 이 채널에서 트랜잭션을 만들거나 트랜잭션을 볼 수 없습니다. 그러나 조직이 채널 생성시에 채널에 피어를 추가 할 수 있습니다.
생성 된 채널에 참여하려면 Participant Org의 콘솔에 액세스하십시오. "Nodes" 탭으로 이동하여 각 피어 노드에 대한 햄버거 메뉴를 클릭 한 다음 "Join New Channels"를 선택하십시오.
채널 이름을 묻는 대화 상자가 표시됩니다. 각 피어에 대해 채널 이름 ( 'samchannel'또는 'judechannel')을 적절하게 입력 한 다음 'Join'을 클릭하십시오.
![](images/join_channel_participant.png)
![](images/join_channel_participant2.png)

다른 모든 peer에 대해서도 동일하게 수행합니다.
* judedealer Org의 콘솔에서도 동일한 작업을 수행합니다.

3. 

## 2. 체인코드 배포하기

## 3. REST API 로 트랜잭션 호출하기

