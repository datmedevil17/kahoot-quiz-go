const API_BASE = 'http://localhost:8080';

export async function api(method: string, path: string, data?: any, token?: string) {
	const headers: HeadersInit = {
		'Content-Type': 'application/json'
	};

	if (token) {
		headers['Authorization'] = `Bearer ${token}`;
	}

	const opts: RequestInit = {
		method,
		headers
	};

	if (data) {
		opts.body = JSON.stringify(data);
	}

	const res = await fetch(`${API_BASE}${path}`, opts);
    
    // Handle non-JSON responses (like 401 Unauthorized sometimes returning string)
    const text = await res.text();
    try {
        return JSON.parse(text); 
    } catch {
        // If not JSON, return struct with error or raw text
        if (!res.ok) return { error: text || res.statusText };
        return { data: text };
    }
}
