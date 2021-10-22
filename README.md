# My SSH Client

쿠버네티스 실습 할 때 서버 접속할때마다 따로 명령어 치기 귀찮아서 만든 클라이언트

## 사용법

### 서버 리스트 출력

```bash
$myssh list -f [파일]
```

### ssh 연결

```bash
#서버 이름으로 연결
$myssh connect -f [파일] -n [이름]

#서버 인덱스 번호로 연결 (list 명령어 쳤을 때 나오는 인덱스)
$myssh connect -f [파일] -i [번호]
```

## 서버리스트 파일 형식

preRun, key 는 선택 사항

```yaml
preRun:
    path: "path/file"
list:
    - name: "sample-01"
      account: "centos"
      ip: "192.168.0.2"
      port: 22
      key: "sampe.pem"
    - name: "sample-02"
      account: "centos"
      ip: "192.168.0.3"
      port: 22
      key: "sampe.pem"

      ...
```
