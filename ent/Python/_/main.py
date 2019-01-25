#! /usr/bin/env python

# =============================================================
# LogicalExpressions - Python
# main.py
# =============================================================
# Created by Lance T. Walker on 1/24/2019
# Copyright (c) 2019 CodeLife-Productions. All rights reserved.
# =============================================================

import os, sys
sys.path.insert(0, os.path.abspath(os.path.join(os.path.dirname(__file__), '..')))

import services as Services
import usecases as Usecases
import problems as Problems
#
# ### Entry ===>
# def main():
#   print "main.py: starting"
# main()

class Main(object):
  def __init__(self, filename):
    self.file = open(filename)

  def __enter__(self):
    self.file.close()

  def __exit__(self, ctx_type, ctx_value, ctx_traceback):
    self.file.close()

with Main('file') as f:
  contents = f.read()
