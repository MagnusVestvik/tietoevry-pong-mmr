import { json } from '@sveltejs/kit';

/**
 * @typedef {Object} Cookies
 * @property {(name: string) => string | undefined} get - Get the value of a cookie.
 * @property {(name: string, value: string, options?: CookieOptions) => void} set - Set a cookie.
 * @property {(name: string, options?: CookieOptions) => void} delete - Delete a cookie.
 */

/**
 * @typedef {Object} CookieOptions
 * @property {string} [path] - The path for which the cookie is valid.
 * @property {Date} [expires] - The expiration date of the cookie.
 * @property {number} [maxAge] - The maximum age of the cookie in seconds.
 * @property {string} [domain] - The domain for which the cookie is valid.
 * @property {boolean} [secure] - Whether the cookie is secure.
 * @property {boolean} [httpOnly] - Whether the cookie is HTTP-only.
 * @property {string} [sameSite] - The SameSite policy for the cookie.
 */

/**
 * Load function to authenticate user and set cookie.
 * 
 * @param {Object} context - The context object.
 * @param {Cookies} context.cookies - The cookies object to manipulate cookies.
 * @returns {Promise<void>} - A promise that resolves when the operation is complete.
 */

export async function POST({ request, cookies }) {
  try {
    const { token } = await request.json();

    if (!token) {
      return json({ error: 'Token is missing' }, { status: 400 });
    }

    cookies.set('Authorization', token, { path: '/' });

    return json({ success: true });
  } catch (error) {
    return json({ error: 'Failed to set cookie' }, { status: 500 });
  }
}

export async function GET({ cookies }) {
  const token = cookies.get('Authorization');

  if (!token) {
    return json({ error: 'Authorization cookie is missing' }, { status: 400 });
  }

  return json({ Authorization: token });
}

export async function DELETE({ cookies }) {
  try {
    cookies.set('Authorization', '', { path: '/', expires: new Date(0) });
    return json({ success: true });
  } catch (error) {
    return json({ error: 'Failed to delete cookie' }, { status: 500 });
  }
}