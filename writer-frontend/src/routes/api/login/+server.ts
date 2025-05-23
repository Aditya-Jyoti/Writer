import { json } from '@sveltejs/kit';
import { env } from '$env/dynamic/private';
import { SignJWT } from 'jose';

const encoder = new TextEncoder();

export async function POST({ request }) {
	const body = await request.json();
	const { username, password } = body;

	const validUsername = env.LOGIN_USERNAME;
	const validPassword = env.LOGIN_PASSWORD;
	const secret = env.JWT_SECRET;

	if (!username || !password || username !== validUsername || password !== validPassword) {
		return json({ error: 'Invalid credentials' }, { status: 401 });
	}

	const token = await new SignJWT({ username })
		.setProtectedHeader({ alg: 'HS256' })
		.setExpirationTime('1d')
		.sign(encoder.encode(secret));

	return json({ token });
}
