using System;
// =============================================================
// LogicalExpressions - CS
// algorithmns.cs
// =============================================================
// Created by Lance T. Walker on 1/22/2019
// Copyright © 2018 CodeLife-Productions. All rights reserved.
// =============================================================

namespace Services {
    class Algorithms {

    }

    interface IAlgorithms<T> {
        void SelectionSort();
        void InsertionSort();
        void BubbleSort();
        void QuickSort();
        void HeapSort();
        void MergeSort();
        void ShellSort(); 
        TopologicalSort(); 
    }
}
