import http from 'k6/http';

export const options = {

  scenarios: {
    basic: {
      executor: 'ramping-vus',
      startTime: '0s',
      gracefulStop: '1s',
      gracefulRampDown: '0s',
      stages: [
        { duration: '10s', target: 40 },
      ]
    },
    high_load: {
      executor: 'ramping-vus',
      startTime: '11s',
      gracefulStop: '1s',
      gracefulRampDown: '0s',
      stages: [
        { duration: '10s', target: 200 },
      ]
    },
    spike: {
      executor: 'ramping-vus',
      startTime: '22s',
      gracefulStop: '1s',
      gracefulRampDown: '0s',
      stages: [
        { duration: '10s', target: 5000 },
      ]
    }
  },

  thresholds: {
    http_req_failed: ['rate < 0.01'],
    http_req_duration: ['p(90) < 100', 'p(95) < 200', 'p(99) < 500'],
  },

  noConnectionReuse: true,

  userAgent: 'MyK6UserAgentString/1.0',

};

export default function () {

  http.get('http://localhost:8099');

}
