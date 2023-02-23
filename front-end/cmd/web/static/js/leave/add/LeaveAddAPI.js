class LeaveAdd_API {
    async addLeave(stringifiedJSON) {
        const url = "http://localhost:8881/leaveAdd";

        const body = {
            method: 'POST',
            body: stringifiedJSON,
        }

        const response = await fetch(url, body)
        const result = await response.json()

        return result
    }
}