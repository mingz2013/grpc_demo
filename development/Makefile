.PHONY: help
help:
	@echo '                                                                          '
	@echo 'Makefile for gitbook doc                                                  '
	@echo '                                                                          '
	@echo 'Usage:                                                                    '
	@echo '   make help                           show help                          '
	@echo '                                                                          '
	@echo '   make up                             启动服务                            '
	@echo '   make down                           停止服务                            '
	@echo '   make logs                           查看日志                            '
	@echo '                                                                          '
	@echo '                                                                          '


.PHONY: up
up:
	docker-compose  up -d


.PHONY: down
down:
	docker-compose  down


.PHONY: logs
logs:
	docker-compose  logs -f


build:
	docker-compose build


scale:
	docker-compose up -d --scale server=10 --no-recreate # 参数，不重新创建，实现滚动更新，扩容缩容不影响访问



etcd:
	etcdctl get --prefix foo