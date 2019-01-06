# Ogame GO bot wrapper
GOlang으로 작성한 오게임 봇 입니다.

##개발환경설정
- 클론 받은 ogameGo_Bot 디렉토리를 GOPATH 로 설정합니다.


##설치 혹은 실행
```go get``` (맥의 경우  gobin 설정이 필요할수 있음 https://stackoverflow.com/a/32357023)


##우선구현기능
- 텔레그램 혹은 discord 연동
- 자동 플릿/성장 명령어 공격

##기본 설계 (아이디어 환영합니다)
- 사람이 브라우저에서 오게임을 즐기는 마냥, TaskQueue 를 두어 순차적으로 Queue에서 Task 를 실행하는 방식
- 1개 바이너리 = 1계정
  - 기본 프레임워크의 캐시 관련 함수의 캐슁 범위 파악 안됨 (``` GetCachedPlanets() []Planet ```)
- 개선 건의사항은 gitlab  내 issue나 기존 카톡방 활용해서 진행하면 좋을것 같습니다.

      
       

##기여자
- Berna
- jc01rho
