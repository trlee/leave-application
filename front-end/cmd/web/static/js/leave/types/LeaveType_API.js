class LeaveType_API {
    // fetch all leave type
    async getAllLeaveType() {
        const url = broker + 'getAllLeaveType'
        const response = await fetch(url)
        const result = await response.json()
        return result
    }
}