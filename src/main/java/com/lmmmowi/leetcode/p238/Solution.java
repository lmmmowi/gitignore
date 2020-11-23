package com.lmmmowi.leetcode.p238;

/**
 * @Author: lmmmowi
 * @Date: 2020/11/23
 * @Description: 238.除自身以外数组的乘积[https://leetcode-cn.com/problems/product-of-array-except-self/]
 */
public class Solution {

    public int[] productExceptSelf(int[] nums) {
        int n = nums.length;
        int[] left = new int[n];
        int[] right = new int[n];
        for (int i = 0; i < n; i++) {
            left[i] = right[i] = 1;
        }

        for (int i = 1; i < n; i++) {
            left[i] = left[i - 1] * nums[i - 1];
        }

        for (int i = n - 2; i >= 0; i--) {
            right[i] = right[i + 1] * nums[i + 1];
        }

        int[] result = new int[n];
        for (int i = 0; i < n; i++) {
            result[i] = left[i] * right[i];
        }
        return result;
    }
}
