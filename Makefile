all: install tests myFind myWc myXargs myRotate

install: clean
	@cd cmd/myFind && go build && cp myFind ../myXargs/
	@cd cmd/myWc && go build && cp myWc ../myXargs/
	@cd cmd/myWc && go build && cp myWc ../myXargs/
	@cd cmd/myXargs && go build
	@cd cmd/myRotate && go build

tests:
	@echo "\e[33mRun tests for myFind:\e[0m"
	@cd internal/find && go test -cover
	@echo ""
	@echo "\e[33mRun tests for myWc:\e[0m"
	@cd internal/count && go test -cover
	@echo ""
	@echo "\e[33mRun tests for myXargs:\e[0m"
	@cd internal/xargs && go test
	@echo ""

clean:
	@cd cmd/myFind && rm -rf myFind
	@cd cmd/myWc && rm -rf myWc
	@cd cmd/myXargs && rm -rf myFind
	@cd cmd/myXargs && rm -rf myWc
	@cd cmd/myXargs && rm -rf myXargs
	@cd cmd/myRotate && rm -rf data/ some_app* myRotate

myFind: install
	@echo ""
	@echo "\e[33mRun command: tree cmd/myFind/foo:\e[0m"
	@cd cmd/myFind/foo && tree
	@echo ""
	@echo "\e[33mRun command: ./myFind foo:\e[0m"
	@cd cmd/myFind && ./myFind foo
	@echo ""
	@echo "\e[33mRun command: ./myFind -f foo:\e[0m"
	@cd cmd/myFind && ./myFind -f foo
	@echo ""
	@echo "\e[33mRun command: ./myFind -d foo:\e[0m"
	@cd cmd/myFind && ./myFind -d foo
	@echo ""
	@echo "\e[33mRun command: ./myFind -sl foo:\e[0m"
	@cd cmd/myFind && ./myFind -sl foo
	@echo ""
	@echo "\e[33mRun command: ./myFind -f -ext "txt" foo:\e[0m"
	@cd cmd/myFind && ./myFind -f -ext "txt" foo

myWc: install
	@echo ""
	@echo "\e[33mRun command: ./myWc input.txt:\e[0m"
	@cd cmd/myWc && ./myWc input.txt
	@echo ""
	@echo ""
	@echo "\e[33mRun command: ./myWc -l input2.txt input3.txt:\e[0m"
	@cd cmd/myWc && ./myWc -l input2.txt input3.txt
	@echo ""
	@echo ""
	@echo "\e[33mRun command: ./myWc -m input.txt input2.txt input3.txt:\e[0m"
	@cd cmd/myWc && ./myWc -m input.txt input2.txt input3.txt
	@echo ""

myXargs: install
	@echo ""
	@echo "\e[33mRun command: tree:\e[0m"
	@cd cmd/myXargs && tree
	@echo ""
	@echo "\e[33mRun command: echo "-la" | ./myXargs ls:\e[0m"
	@cd cmd/myXargs && echo "-la" | ./myXargs ls
	@echo ""
	@echo "\e[33mRun command: ./myFind -f -ext 'txt' somedir | ./myXargs ./myWc -m:\e[0m"
	@cd cmd/myXargs && ./myFind -f -ext 'txt' somedir | ./myXargs ./myWc -m
	@echo ""

myRotate: install
	@echo ""
	@echo "\e[33mRun command:./myRotate /logs/some_app.log /logs/some_app2.log /logs/some_app3.log && tree:\e[0m"
	@cd cmd/myRotate && ./myRotate logs/some_app.log logs/some_app2.log logs/some_app3.log && tree
	@echo ""
	@echo "\e[33mRun command:./myRotate -a data/archive /logs/some_app.log /logs/some_app2.log /logs/some_app3.log && tree:\e[0m"
	@cd cmd/myRotate && ./myRotate -a data/archive logs/some_app.log logs/some_app2.log logs/some_app3.log && tree
	@echo ""