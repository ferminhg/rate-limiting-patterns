import http from 'k6/http';
import { check, sleep } from 'k6';

export let options = {
    stages: [
        { duration: '1m', target: 200 },
        { duration: '2m', target: 300 },
        { duration: '2m', target: 400 },
        { duration: '1m', target: 0 },
    ],
};

export default function () {
    // success
    let res = http.get('http://leaky-bucket:3010/');
    check(res, {
        'status is 200': (r) => r.status === 200,
        'response is Success': (r) => r.body === 'Success',
    });

    // Optional: Add a sleep to simulate user think time
    sleep(1);
}