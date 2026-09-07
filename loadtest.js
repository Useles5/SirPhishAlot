//------ AI Generated [O_O] Load test script ------//

import http from 'k6/http';
import { check } from 'k6';
import { SharedArray } from 'k6/data';

const corpus = new SharedArray('certs', function () {
    // Read the file, split by newline, and drop any empty trailing lines
    return open('./data.jsonl').split('\n').filter(line => line.trim().length > 0);
});

// 2. Configure the stress test
export const options = {
    vus: 100,           // 100 concurrent virtual users
    duration: '15s',    // 15-second blast to get stable metrics
};

const params = {
    headers: { 'Content-Type': 'application/json' },
};

// 3. The attack loop
export default function () {
    // Pick a payload based on the Virtual User ID and iteration count.
    // This ensures VUs cycle through the 50,000 messy records dynamically
    // without the CPU branch predictor memorizing the payload shape.
    const payload = corpus[(__VU * 1000 + __ITER) % corpus.length];

    // Fire the POST request at your Go server
    const res = http.post('http://localhost:8080/ingest', payload, params);

    // Verify the Go server didn't choke on a malformed payload
    check(res, {
        'status is 200': (r) => r.status === 200,
    });
}