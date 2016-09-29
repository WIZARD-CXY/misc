KVERS = $(shell uname -r)

obj-m += killd.o

build: kernel_modules

kernel_modules:
        make -C /usr/src/linux-headers-$(KVERS) M=$(CURDIR) modules
clean:
        make -C /lib/modules/$(KVERS)/build M=$(CURDIR) clean
