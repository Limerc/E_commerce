.PHONY: gen-demo-proto
gen-demo-proto:
	@cd demo/demo_proto && cwgo server -I ../../idl --type RPC --module github.com/Limerc/E_commerce/gomall/demo/demo_proto --service demo_proto --idl ../../idl/echo.proto

.PHONY: gen-demo-thrift
gen-demo-thrift:
	@cd demo/demo_thrift && cwgo server --type RPC --module github.com/Limerc/E_commerce/gomall/demo/demo_thrift --service demo_thrift --idl ../../idl/echo.thrift

.PHONY: gen-frontend
gen-frontend:
	@cd app/frontend && cwgo server --type HTTP --idl ../../idl/frontend/auth_page.proto --service frontend -module github.com/Limerc/E_commerce/gomall/app/frontend -I ../../idl