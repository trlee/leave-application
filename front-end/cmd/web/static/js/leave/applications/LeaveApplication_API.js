const broker = "http://localhost:8881/"

class LeaveApplication_API {
    // apply a new leave
    async applyLeave(stringifiedJSON) {
        const url = broker + "leaveApply";

        const body = {
            method: 'POST',
            body: stringifiedJSON,
        }

        const response = await fetch(url, body)
        const result = await response.json()

        return result
    }

    // fetch all leave application
    async getAllLeaveApplication() {
        const url = broker + 'getAllLeaveApplication';
        const response = await fetch(url);
        const result = await response.json();
        return result;
  }
}