import { writable } from "svelte/store";

export const page_idx = writable(1);
// 用一个 store 来保存当前主题状态（true = 深色 / synthwave，false = 浅色 / 默认 light）
export const isDark = writable(false);