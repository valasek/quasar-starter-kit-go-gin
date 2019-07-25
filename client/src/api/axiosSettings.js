// Copyright © 2019 Stanislav Valasek <valasek@gmail.com>

import axios from 'axios'

var baseURL = ''

const debug = process.env.NODE_ENV !== 'production'
if (debug) {
  baseURL = 'http://localhost:3000'
}

const apiClient = axios.create({
  baseURL: baseURL,
  withCredentials: false, // This is the default
  crossDomain: true,
  headers: {
    Accept: 'application/json',
    'Content-Type': 'application/json'
  },
  timeout: 10000
})

export default {
  apiClient
}
