class LeaveApply_API {
    async applyLeave(stringifiedJSON) {
        const url = "http://localhost:8881/leave/apply";

        const body = {
            method: 'POST',
            body: stringifiedJSON,
        }

        const response = await fetch(url, body)
        const result = await response.json()

        return result
    }
}