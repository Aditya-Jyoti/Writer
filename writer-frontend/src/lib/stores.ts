import { writable } from 'svelte/store';

import type { Blog } from '../app.d.ts';

export const allBlogs = writable<Blog[] | []>([]);
