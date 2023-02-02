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
@Table(name = "node_type")
public class NodeTypeEntity {

    @Id
    @Column(name = "node_type_cd")
    private String nodeTypeCd;
    @Column(name = "description")
    private String description;
    @Column(name = "active_bit")
    private Integer activeBit;

    @OneToMany(mappedBy = "nodeTypeEntity")
    private List<ObjectRefEntity> objectRefEntityList;

}
