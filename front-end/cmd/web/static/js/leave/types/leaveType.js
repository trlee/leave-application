const   TypeHelpers = new LeaveType_Helpers(),
        TypeAPI = new LeaveType_API(),
        dt = 'leaveType',
        dtBody = 'leaveTypeBody'

window.addEventListener('DOMContentLoaded', () => {
    // fetch all leave types & update DOM
    TypeAPI.getAllLeaveType().then(resp => {
        console.log(resp)
        TypeHelpers.insertRows(dtBody, resp.data)
        TypeHelpers.triggerDT(dt, dtBody)
    })
})