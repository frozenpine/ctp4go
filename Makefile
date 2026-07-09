VERSION ?= v6.7.13
PLATFORM ?= future
GEN_TYPES ?= 0

SUB_MODULES = thost trader mduser

CLEAN_SUBMODULES = $(addprefix clean-, $(SUB_MODULES))

export VERSION PLATFORM

all: $(SUB_MODULES)

$(SUB_MODULES):
	$(MAKE) -C $@

$(CLEAN_SUBMODULES):
	$(MAKE) -C $(patsubst clean-%,%,$@) clean

clean: $(CLEAN_SUBMODULES)

.PHONY: all clean $(SUB_MODULES) $(CLEAN_SUBMODULES)