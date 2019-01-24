#! /usr/bin/env python

# =============================================================
# LogicalExpressions - Python
# main.py
# =============================================================
# Created by Lance T. Walker on 1/24/2019
# Copyright (c) 2019 CodeLife-Productions. All rights reserved.
# =============================================================

import os, sys
modules = ['./usecases', './services', './problems']
for path in modules:
  sys.path.insert(0, path)
import algorithms
import checks, conversions
import project_euler

# Entry ===>
def main():
  print "main.py: Starting"

if __name__ == "__main__":
  main()
