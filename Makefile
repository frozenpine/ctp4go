VERSION ?= v6.7.13
PLATFORM ?= future

SUB_MODULES = thost trader mduser

CLEAN_SUBMODULES = $(addprefix clean-, $(SUB_MODULES))

export VERSION PLATFORM

.PHONY: all clean $(SUB_MODULES) $(CLEAN_SUBMODULES)

all: $(SUB_MODULES)

$(SUB_MODULES):
	$(MAKE) -C $@

$(CLEAN_SUBMODULES):
	$(MAKE) -C $(patsubst clean-%,%,$@) clean

clean: $(CLEAN_SUBMODULES)