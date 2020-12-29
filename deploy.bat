set GOOS=linux
set GOARCH=arm
set GOARM=5

set KINDLE=192.168.1.65

go build sample/main.go
move main app/gosample

ssh root@%KINDLE% rm -rf /mnt/us/extensions/gosample
scp -r app root@%KINDLE%:/mnt/us/extensions/gosample