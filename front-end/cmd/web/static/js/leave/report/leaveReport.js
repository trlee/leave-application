const ReportHelpers = new LeaveReport_Helpers(),
      ReportAPI = new LeaveReport_API(),
      dt = 'leaveReport',
      dtBody = 'leaveTypeBody'

window.addEventListener('DOMContentLoaded', () => {
    // fetch all leave types & update DOM
    ReportAPI.getAllLeaveReport().then(resp => {
        console.log(resp)
        TypeHelpers.insertRows(dtBody, resp.data)
        TypeHelpers.triggerDT(dt, dtBody)
    })
})