# 2020.7.31 최광훈 박사님 하이퍼렛져 수업

# 베이직 네트워크 구성
## - 1.준비물 생성단계
    crypto-config.yaml
    configtx.yaml
    generate.sh
    crypto-config dir
    config/ genesis.block, ch.tx

## - 2.네트워크 구성
    .env
    docker-compose -> up
    ca, peer (x3) , couchDb(x3), ord erer, cli
    net_basic

## - 3.채널 구성
    cli ch.tx
    peer channel create
    mychannel.block
    각 피어별로 peer channel join

## - 4.체인 코드 작성
    Go, java, nodejs
    shim
    peer
    contract
    설치 -> 피어
    배포 -> 채널
    테스트
    
## - 실습 고고
    디렉토리 구성-네트워크,컨트랙트,어플리케이션

## 오늘의 결과

- [x] 1.준비물 생성단계
- [x] 2.네트워크 구성
- [x] 3.채널 구성
- [x] 4.체인코드 작성
- [x] 5.체인코드 배포
- [ ] 6.teamate 구현


## 

-----------------------------

# 2020.8.3 최광훈 박사님 하이퍼렛져 수업

## - 체인코드 작성 및 쉘 스크립트 작성

### - 1. 체인코드 작성 및 컴파일링
    sacc.go 를 카피해서 나의 디렉토리로 옴겼다.
    그다음 go build 를 하기위해
    go get -u "github.com/hyperledger/fabric/core/chaincode/shim" 를 하고

    go build 를 진행함.
    이걸 하는이유는 go로 만든 체인코드를 컴퓨터가 이해할수있는 언어로 컴파일
    해야하기 때문이다.

### - 2. 컴파일링 된 체인코드를 쉘 스크립트를 통해 배포
    실행 전 ./teardown.sh 명령어를 통해 네트워크를 초기화 시키고 난 후 다시 ./start.sh 명령어를 통해 다시 네트워크 구축

    그다음 ./cc_tea.sh 쉘 스크립트 를 작성
    쉘 스크립트 작성할때 instatiate 및 upgrade 를
    해봤다. 정상적으로 작동 됨
    