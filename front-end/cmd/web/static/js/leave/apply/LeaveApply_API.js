const broker = "http://localhost:8881/"

class LeaveApply_API {
    // apply a new leave
    async applyLeave(stringifiedJSON) {
        const url = broker + 'leaveApply';

        const body = {
            method: 'POST',
            body: stringifiedJSON,
        }

        const response = await fetch(url, body)
        const result = await response.json()

        return result
    }
}