package com.crossasyst.activities.entity;

import lombok.AllArgsConstructor;
import lombok.Data;
import lombok.NoArgsConstructor;

import javax.persistence.*;
import java.util.List;

@Data
@AllArgsConstructor
@NoArgsConstructor
@Entity
@Table(name = "activity_type")
public class ActivityTypeEntity {

    @Id
    @Column(name = "activity_type_cd")
    private String activityTypeCd;
    @Column(name = "description")
    private String description;
    @Column(name = "active_bit")
    private Integer activeBit;

    @OneToMany(mappedBy = "activityTypeEntity")
    private List<ActivityEntity>activityEntitiesList;

}
