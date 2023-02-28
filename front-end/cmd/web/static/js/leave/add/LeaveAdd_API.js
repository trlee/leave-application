broker = "http:/localhost:8881/"

class LeaveAdd_API {
    // add a new leave type
    async addLeave(stringifiedJSON) {
        const url = broker + 'leaveTypeAdd'

        const body = {
            method: 'POST',
            body: stringifiedJSON,
        }

        const response = await fetch(url, body)
        const result = await response.json()

        return result
    }
}