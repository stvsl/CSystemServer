# Additional clean files
cmake_minimum_required(VERSION 3.16)

if("${CONFIG}" STREQUAL "" OR "${CONFIG}" STREQUAL "Debug")
  file(REMOVE_RECURSE
  "CMakeFiles/CSystem-Qt6_autogen.dir/AutogenUsed.txt"
  "CMakeFiles/CSystem-Qt6_autogen.dir/ParseCache.txt"
  "CSystem-Qt6_autogen"
  )
endif()
