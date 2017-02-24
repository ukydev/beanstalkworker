package beanstalkworker

import "time"
import "github.com/kr/beanstalk"
import "fmt"
import "log"

// RawJob represents the raw job data that is returned by beanstalkd.
type RawJob struct {
	id          uint64
	err         error
	body        *[]byte
	conn        *beanstalk.Conn
	tube        string
	prio        uint32
	releases    uint32
	age         time.Duration
	returnPrio  uint32
	returnDelay time.Duration
}

// Delete function deletes the job from the queue.
func (job *RawJob) Delete() {
	if err := job.conn.Delete(job.id); err != nil {
		job.LogError("Could not delete job: " + err.Error())
	}
}

// Release function releases the job from the queue.
func (job *RawJob) Release() {
	if err := job.conn.Release(job.id, job.returnPrio, job.returnDelay); err != nil {
		job.LogError("Could not release job: " + err.Error())
	}
}

// Bury function buries the job from the queue.
func (job *RawJob) Bury() {
	if err := job.conn.Bury(job.id, job.returnPrio); err != nil {
		job.LogError("Could not bury job: " + err.Error())
	}
}

// SetReturnPriority sets the return priority to use if a job is released or buried.
func (job *RawJob) SetReturnPriority(prio uint32) {
	job.returnPrio = prio
}

// SetReturnDelay sets the return delay to use if a job is released back to queue.
func (job *RawJob) SetReturnDelay(delay time.Duration) {
	job.returnDelay = delay
}

// GetAge gets the age of the job from the job stats.
func (job *RawJob) GetAge() time.Duration {
	return job.age
}

// GetPriority gets the priority of the job.
func (job *RawJob) GetPriority() uint32 {
	return job.prio
}

// GetReleases gets the count of release of the job.
func (job *RawJob) GetReleases() uint32 {
	return job.releases
}

// GetTube returns the tube name we got this job from.
func (job *RawJob) GetTube() string {
	return job.tube
}

// LogError function logs an error messagge regarding the job.
func (job *RawJob) LogError(a ...interface{}) {
	log.Print("Tube: ", job.tube, ", Job: ", job.id, ": Error: ", fmt.Sprint(a...))
}

// LogInfo function logs an info messagge regarding the job.
func (job *RawJob) LogInfo(a ...interface{}) {
	log.Print("Tube: ", job.tube, ", Job: ", job.id, ": ", fmt.Sprint(a...))
}