import type { PageLoad } from './$types';

/**
 * Fetches all blog posts from the backend API.
 *
 * @summary Get all blog posts
 * @description Retrieves all blog entries from the database via the writer-backend service.
 * @route GET /blogs/
 * @returns {Promise<{ blogs: Blog[] }>} - A promise resolving to an object containing the list of blog posts.
 * @throws Will throw an error if the fetch request fails or returns a non-OK status.
 */
export const load: PageLoad = async ({ fetch }) => {
	const res = await fetch(`${import.meta.env.PUBLIC_API_BASE}/blogs`);
	if (!res.ok) {
		throw new Error('Failed to fetch blogs');
	}

	const blogs = await res.json();
	return { blogs };
};
