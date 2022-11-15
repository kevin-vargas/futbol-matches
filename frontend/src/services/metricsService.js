import config from '../config/config'

const metricsService = {
    getMetric: (metricType, interval) => {
        const url = `${config.apiHost}${config.endpoints.getMetrics.path}/${metricType}?query=${interval}`
        return fetch(url);
    }
}

export default metricsService;
