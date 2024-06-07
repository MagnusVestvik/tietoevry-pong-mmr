export async function load({ cookies }) {
	
	try {
		const response = await fetch("http://localhost:8080/api/login", {
		  method: 'POST',
		  headers: {
			'Content-Type': 'application/json'
		  },
		  body: JSON.stringify({ email: 'm@picheta.dev', password: 'longpassword1!' }) // Replace with actual credentials
		});
  
		if (!response.ok) {
		  throw new Error('Login failed');
		}
  
		const data = await response.json();
		cookies.set('Authorization', data.token, { path: '/' });
  
	  } catch (error) {
		console.error('Error logging in:', error);
	}
}