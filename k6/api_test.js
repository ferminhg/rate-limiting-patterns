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

const RATE_LIMITER_URL = 'http://rate-limiter-service:3010';

export default function () {
    makeNoLimiterRequest();
    makeLeakyBucketRequest();
}

function makeNoLimiterRequest() {
    let res = http.get(RATE_LIMITER_URL + '/no-limiter');
    check(res, {
        'status is 200': (r) => r.status === 200,
        'response is Success': (r) => r.body === 'Success',
    });
    sleep(1);
}

function makeLeakyBucketRequest() {
    let res = http.get(RATE_LIMITER_URL + '/leaky-bucket');
    check(res, {
        'status is 200': (r) => r.status === 200,
        'response is Success': (r) => r.body === 'Success',
    });
    sleep(1);
}
