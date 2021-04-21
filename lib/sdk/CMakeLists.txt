cmake_minimum_required(VERSION 3.15)
project(sdk C)

include(CheckIncludeFile)

set(CMAKE_C_STANDARD 99)
set(SAMP_SDK_ROOT ${PROJECT_SOURCE_DIR})

add_library(sdk STATIC ${SAMP_SDK_ROOT}/amxplugin.c ${SAMP_SDK_ROOT}/amxplugin2.c ${SAMP_SDK_ROOT}/amx/getch.c)
target_include_directories(sdk PUBLIC ${SAMP_SDK_ROOT} ${SAMP_SDK_ROOT}/amx)

check_include_file(alloca.h HAVE_ALLOCA_H)
if(HAVE_ALLOCA_H)
	target_compile_definitions(sdk PUBLIC -DHAVE_ALLOCA_H)
endif()

check_include_file(inttypes.h HAVE_INTTYPES_H)
if(HAVE_INTTYPES_H)
	target_compile_definitions(sdk PUBLIC -DHAVE_INTTYPES_H)
endif()

check_include_file(malloc.h HAVE_MALLOC_H)
if(HAVE_MALLOC_H)
	target_compile_definitions(sdk PUBLIC -DHAVE_MALLOC_H)
endif()

check_include_file(stdint.h HAVE_STDINT_H)
if(HAVE_STDINT_H)
	target_compile_definitions(sdk PUBLIC -DHAVE_STDINT_H)
endif()

if (WIN32 AND MSVC)
	target_compile_definitions(sdk PUBLIC -D_CRT_SECURE_NO_WARNINGS -DWIN32)
	message(STATUS "On windows with msvc")
elseif(UNIX)
	target_compile_definitions(sdk PUBLIC -DLINUX)
	target_compile_options(sdk PRIVATE -m32 -std=c99)
	
	message(STATUS "On unix")
else()
	target_compile_definitions(sdk PRIVATE -DWIN32)
	target_compile_options(sdk PRIVATE -m32 -std=c99)

	if (CYGWIN)
		# Using cygwin compiler

		target_compile_options(sdk PUBLIC -Wno-ignored-attributes)
		message(STATUS "Using cygwin")
	elseif(MINGW)
		# Using mingw
		# Some really stupid but necessary workaround

		target_compile_options(sdk PUBLIC -Wno-attributes -include stddef.h)
		message(STATUS "Using mingw")
	else()
		# Some unsupported platform

		message(FATAL_ERROR "This platform is unsupported")
	endif()
endif()

set_target_properties(sdk PROPERTIES LINKER_LANGUAGE C)
message(STATUS "Completed")
