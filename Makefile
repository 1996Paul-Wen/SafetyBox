# refer to https://seisman.github.io/how-to-write-makefile/overview.html

GOCC = go
TARGET = bin/safetybox
M_SRC = app/*.go

# the first target is the default target of Makefile
# target `ultimate_target` requires target `prepare` and target $(TARGET)
.PHONY: ultimate_target
ultimate_target:  prepare  $(TARGET)  

.PHONY: prepare
prepare:
# `prepare` has no dependencies, so the following command are run
	@echo "working on target [prepare]"
	@mkdir -p bin/

# .PHONY后面跟的目标都被称为伪目标，也就是说我们 make 命令后面跟的参数如果出现在.PHONY 定义的伪目标中，那就直接在Makefile中就执行伪目标的依赖和命令。
# 不管Makefile同级目录下是否有该伪目标同名的文件
.PHONY: $(TARGET)
$(TARGET): $(M_SRC)
	@echo "working on building $(TARGET)"
	@echo "--- building $(M_SRC):"
	$(GOCC) build -o $@ $^

.PHONY : clean
clean:
	@echo "cleaning $(TARGET)"
	-@rm -rf $(TARGET)
	@echo "cleaned"