


@echo on
for %%i in ( *.proto ) do (
protoc --proto_path=./ --micro_out=./proto/ --go_out=./proto/ %%i
)
pause