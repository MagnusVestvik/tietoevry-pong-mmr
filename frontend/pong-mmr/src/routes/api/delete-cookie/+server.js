import { json } from '@sveltejs/kit';

export async function GET({ cookies }) {
	if (!cookies)
		return json({ error: 'Cookie is missing' }, { status: 400 });
	return json({ Authorization: cookies });
}
