git fetch
git checkout %GIT_COMMIT%

powershell -command set-executionpolicy remotesigned
powershell .\bin\replace-sha.ps1

SET GOPATH=%CD%\Godeps\_workspace;c:\Users\Administrator\go
c:\Go\bin\go test -v .
