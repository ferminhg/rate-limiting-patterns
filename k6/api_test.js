import http from 'k6/http';
import { check, sleep } from 'k6';

export let options = {
    stages: [
        { duration: '1m', target: 10 }, // Ramp-up to 10 users over 1 minute
        { duration: '2m', target: 10 }, // Stay at 10 users for 2 minutes
        { duration: '1m', target: 0 },  // Ramp-down to 0 users over 1 minute
    ],
};

export default function () {
    sleep(5);
    // success
    let res = http.get('http://leaky-bucket:3010/');
    check(res, {
        'status is 200': (r) => r.status === 200,
        'response is Success': (r) => r.body === 'Success',
    });

    // Optional: Add a sleep to simulate user think time
    sleep(1);
}