gen:
	@go generate

build: gen
	@go build -o builds/epbf_drop

run: 
	@sudo ./builds/epbf_drop 

clean:
	@rm -rf builds tcp_drop_bpf*
