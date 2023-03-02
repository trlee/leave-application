class LeaveReport_API {
    // fetch all leave report
    async getAllLeaveReport() {
        const url = broker + 'getAllLeaveReport'
        const response = await fetch(url)
        const result = await response.json()
        return result
    }
}