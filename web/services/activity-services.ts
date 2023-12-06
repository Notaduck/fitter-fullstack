import axios from "axios";

export type Record = {
  CreatedAt: string;
  UpdatedAt: string;
  DeletedAt: string | null;
  ID: number;
  ActivityID: number;
  AccumulatedPower: number;
  ActivityType: number;
  Altitude: number;
  BallSpeed: number;
  Cadence: number;
  Cadence256: number;
  Calories: number;
  CombinedPedalSmoothness: number;
  CompressedAccumulatedPower: number;
  CompressedSpeedDistance: string;
  CycleLength: number;
  Cycles: number;
  DeviceIndex: number;
  Distance: number;
  EnhancedAltitude: number;
  EnhancedSpeed: number;
  FractionalCadence: number;
  GPSAccuracy: number;
  Grade: number;
  HeartRate: number;
  LeftPedalSmoothness: number;
  LeftRightBalance: number;
  LeftTorqueEffectiveness: number;
  PositionLat: number;
  PositionLong: number;
  Power: number;
  Resistance: number;
  RightPedalSmoothness: number;
  RightTorqueEffectiveness: number;
  SaturatedHemoglobinPercent: number;
  SaturatedHemoglobinPercentMax: number;
  SaturatedHemoglobinPercentMin: number;
  Speed: number;
  Speed1s: string;
  StanceTime: number;
  StanceTimePercent: number;
  StrokeType: number;
  Temperature: number;
  Time128: number;
  TimeFromCourse: number;
  Timestamp: string;
  TotalCycles: number;
  TotalHemoglobinConc: number;
  TotalHemoglobinConcMax: number;
  TotalHemoglobinConcMin: number;
  VerticalOscillation: number;
  VerticalSpeed: number;
  Zone: number;
  ActivityId: number;
}

export type Activity = {
  ID: number;
  CreatedAt: string;
  UpdatedAt: string;
  DeletedAt: null | string;
  id: number;
  distance: number;
  elevation: number;
  totalRideTime: number;
  timeStamp: string;
  totalTimerTime: number;
  numSessions: number;
  type: number;
  event: number;
  eventType: number;
  localTimestamp: string;
  eventGroup: number;
  userId: number;
  records: Record[];
};
class ActivityService {

  //   /**
  //    * Get all activitites 
  //    * @returns
  //    */
  getAllActivities() {
    const response = axios.get<Activity[]>('http://localhost:3000/api/activities').then(e => e.data)
    return response
  }


  //   /**
  //    * Get By Id
  //    * @returns
  //    */
  getActivityById(id: number) {
    return axios.get<Activity>(`http://localhost:3000/api/activity?activityId=${id}`).then(res => res.data);
  }


  //   /**
  //    *Add activity
  //    * @returns
  //    */
  async addActivity(activities) {
    const res = await axios.post("http://localhost3000/api/activity", {
      method: "POST",
      data: activities,
      headers: {
        "Content-Type": "multipart/form-data"
      },
    });
    return res;
  }

  //   /**
  //    *To Update a Post
  //    * @returns
  //    */
  //   async updatePost() {
  //     const res = await axios.put(
  //       "https://jsonplaceholder.typicode.com/posts/1",
  //       {
  //         method: "PUT",
  //         body: JSON.stringify({
  //           id: 1,
  //           title: "foo",
  //           body: "bar",
  //           userId: 1,
  //         }),
  //         headers: {
  //           "Content-type": "application/json; charset=UTF-8",
  //         },
  //       }
  //     );
  //     return res;
  //   }

  //   /**
  //    *To Delete a Post
  //    * @returns
  //    */
  //   async deletePost() {
  //     const res = await axios.delete(
  //       "https://jsonplaceholder.typicode.com/posts/1",
  //       {
  //         method: "DELETE",
  //       }
  //     );
  //     return res;
  //   }
}

export default new ActivityService();