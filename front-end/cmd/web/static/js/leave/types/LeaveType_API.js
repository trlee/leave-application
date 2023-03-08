class LeaveType_API {
    // fetch all leave type
    async getAllLeaveType() {
        const url = 'http://localhost:8881/getAllLeaveType'
        const response = await fetch(url)
        console.log('getLeaveType here...', url)
        const result = await response.json()
        return result
    }

    async updateLeaveType(stringifiedJSON) {
        const url = broker + 'leaveTypeUpdate'
        const body = {
            method: 'POST',
            body: stringifiedJSON,
        }
        const response = await fetch(url, body)
        const result = await response.json()
        return result
    }

    async deleteLeaveType(ID) {
        const url = broker + 'leaveTypeDelete'
        const body = {
            method: 'POST',
            body: JSON.stringify(ID),
        }
        const response = await fetch(url, body)
        const result = await response.json()
        return result
    }
}