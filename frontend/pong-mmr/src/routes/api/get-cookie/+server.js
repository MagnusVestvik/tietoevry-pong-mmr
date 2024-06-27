import { json } from '@sveltejs/kit';


export async function GET({ cookies }) {
  const token = cookies.get('Authorization');

  if (!token) {
    return json({ error: 'Authorization cookie is missing' }, { status: 400 });
  }

  return json({ Authorization: token });
}
